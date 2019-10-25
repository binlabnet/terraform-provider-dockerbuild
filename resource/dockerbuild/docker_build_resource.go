package dockerbuild

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"regexp"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
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
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCreate(d *schema.ResourceData, m interface{}) error {
	dc, err := client.NewEnvClient()
	if err != nil {
		return errors.Wrap(err, "while creating docker client")
	}

	sourceHash := d.Get("source_hash").(string)
	sourceDir := d.Get("source_dir").(string)

	imageID := fmt.Sprintf("sourcebuild:%s", sourceHash)

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
		Tags: []string{imageID},
	})

	if err != nil {
		return errors.Wrap(err, "while creating tar uploader")
	}

	defer ibResponse.Body.Close()
	dec := json.NewDecoder(ibResponse.Body)

	builtRegexp, err := regexp.Compile("^Successfully built ([0-9a-z]+)+\n$")

	if err != nil {
		return errors.Wrap(err, "while compiling regexp")
	}

	lastID := ""

	for {
		msg := jsonmessage.JSONMessage{}
		err = dec.Decode(&msg)
		if err == io.EOF {
			break
		}

		if err != nil {
			return errors.Wrap(err, "while reading build output")
		}

		matches := builtRegexp.FindStringSubmatch(msg.Stream)

		if len(matches) > 1 {
			lastID = matches[1]
		}

		if msg.Error != nil {
			return errors.Errorf("error code %d building image: %s", msg.Error.Code, msg.Error.Message)
		}
	}

	ii, _, err := dc.ImageInspectWithRaw(context.Background(), lastID)
	if err != nil {
		return errors.Wrapf(err, "while getting image with id %s", lastID)
	}

	err = d.Set("image_id", ii.ID)
	if err != nil {
		return errors.Wrap(err, "while setting image_id")
	}

	d.SetId(sourceHash)

	return resourceRead(d, m)
}

func resourceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceRead(d, m)
}

func resourceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
