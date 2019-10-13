package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.com/noname-ltd/terraform-provider-dockerbuild/datasource/gittree"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"dockerbuild_git_tree": gittree.Resource(),
		},
		ResourcesMap: map[string]*schema.Resource{},
	}
}
