package gittree

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Read: resourceRead,

		Schema: map[string]*schema.Schema{
			"git_root": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source_dir": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tree_sha": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceRead(d *schema.ResourceData, m interface{}) error {
	gitDir := d.Get("git_root").(string)

	repo, err := git.PlainOpen(gitDir)
	if err != nil {
		return errors.Wrapf(err, "while opening git dir %s", gitDir)
	}

	head, err := repo.Head()
	if err != nil {
		return errors.Wrap(err, "while getting git head")
	}

	commit, err := repo.CommitObject(head.Hash())
	if err != nil {
		return errors.Wrap(err, "while getting head commit")
	}

	root, err := commit.Tree()
	if err != nil {
		return errors.Wrap(err, "while getting root tree")
	}

	sourceDir := d.Get("source_dir").(string)

	ent, err := root.FindEntry(sourceDir)
	if err != nil {
		return errors.Wrapf(err, "while finding entry %s", sourceDir)
	}

	err = d.Set("tree_sha", ent.Hash.String())
	if err != nil {
		return errors.Wrap(err, "while setting tree_sha")
	}

	d.SetId(sourceDir)

	return nil
}

// func MD5All(root string) (map[string][md5.Size]byte, error) {
// 	m := make(map[string][md5.Size]byte)
// 	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		if !info.Mode().IsRegular() {
// 			return nil
// 		}
// 		data, err := ioutil.ReadFile(path)
// 		if err != nil {
// 			return err
// 		}
// 		m[path] = md5.Sum(data)
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return m, nil
// }
