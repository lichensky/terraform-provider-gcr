# terraform-provider-gcr

Terraform provider for fetching images digest from Google Container Registry.

## Usage

Example usage:

```
provider "gcr" {
    project = "my-project"
}

data "gcr_image" "my_image" {
    name = "my-image"
    tag  = "master"
}

module "deployment" {
    ...
    image_digest = data.gcr_image.my_image.digest
}
```

## Installation

1. Download provider from *Releases* tab
2. Unzip the archive and move `terraform-provider-gcr` to Terraform's plugin
   directory (see [docs](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins)).

## Authentication

Provider uses `GOOGLE_APPLICATION_CREDENTIALS` environment variable to
authenticate with GCR.
