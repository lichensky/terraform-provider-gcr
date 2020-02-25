package gcr

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// Config representation
type Config struct {
	Project string
}

// Provider return a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"gcr_image": dataSourceGCRImage(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Project: d.Get("project").(string),
	}

	return config, nil
}
