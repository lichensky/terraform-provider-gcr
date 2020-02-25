package gcr

import (
	"fmt"

	gcrname "github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/google"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const gcrDomain = "gcr.io"

func dataSourceGCRImage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGCRImageRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tag": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "latest",
			},
			"digest": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGCRImageRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	name := d.Get("name").(string)
	tag := d.Get("tag").(string)

	repositoryName := fmt.Sprintf("%s/%s/%s", gcrDomain, config.Project, name)

	repo, err := gcrname.NewRepository(repositoryName)
	if err != nil {
		return err
	}

	auth, err := google.NewEnvAuthenticator()
	if err != nil {
		return err
	}

	reference := repo.Tag(tag)
	image, err := remote.Image(reference, remote.WithAuth(auth))
	if err != nil {
		return err
	}

	digest, err := image.Digest()
	if err != nil {
		return err
	}

	d.Set("digest", digest.String())
	d.SetId(digest.String())

	return nil
}
