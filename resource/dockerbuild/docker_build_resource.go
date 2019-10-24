package dockerbuild

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"source_dir": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source_hash": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCreate(d *schema.ResourceData, m interface{}) error {
	dc, err := client.NewEnvClient()
	if err != nil {
		return errors.Wrap(err, "while creating docker client")
	}

	imageName := d.Get("image_name").(string)
	sourceHash := d.Get("source_hash").(string)
	sourceDir := d.Get("source_dir").(string)

	imageID := fmt.Sprintf("%s:%s", imageName, sourceHash)

	_, _, err = dc.ImageInspectWithRaw(context.Background(), imageID)

	if !client.IsErrNotFound(err) {
		return resourceRead(d, m)
	}

	rc, err := archive.Tar(sourceDir, archive.Uncompressed)
	defer rc.Close()

	if err != nil {
		return errors.Wrap(err, "while creating tar uploader")
	}

	ibResponse, err := dc.ImageBuild(context.Background(), rc, types.ImageBuildOptions{
		Tags: []string{"whatever:oldest"},
	})

	if err != nil {
		return errors.Wrap(err, "while creating tar uploader")
	}

	out, err := ioutil.ReadAll(ibResponse.Body)
	if err != nil {
		return errors.Wrap(err, "while reading build response")
	}

	ioutil.WriteFile("/tmp/output", out, 0700)

	d.SetId(imageID)

	return resourceRead(d, m)
}

func resourceRead(d *schema.ResourceData, m interface{}) error {
	dc, err := client.NewEnvClient()
	if err != nil {
		return errors.Wrap(err, "while creating docker client")
	}

	imageName := d.Get("image_name").(string)
	sourceHash := d.Get("source_hash").(string)

	imageID := fmt.Sprintf("%s:%s", imageName, sourceHash)

	_, _, err = dc.ImageInspectWithRaw(context.Background(), imageID)

	if !client.IsErrNotFound(err) {
		d.SetId("")
	}

	return nil
}

func resourceUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceRead(d, m)
}

func resourceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
