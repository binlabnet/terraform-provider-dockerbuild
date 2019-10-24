package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.com/noname-ltd/terraform-provider-dockerbuild/datasource/sourcetree"
	"gitlab.com/noname-ltd/terraform-provider-dockerbuild/resource/dockerbuild"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"dockerbuild_source_tree": sourcetree.Resource(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"dockerbuild_build": dockerbuild.Resource(),
		},
	}
}
