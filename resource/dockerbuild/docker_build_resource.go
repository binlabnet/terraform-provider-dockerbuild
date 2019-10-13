package dockerbuild

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCreate(d *schema.ResourceData, m interface{}) error {
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
