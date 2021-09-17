# terraform-provider-workspace
A provider that errors out if the current terraform workspace does not match the input variables


## Usage
Add the provider to your list of required providers and use `terraform init`:

```hcl
terraform {
  required_providers {
    workspace = {
      source = "erran/workspace"
      version = ">= 0.0.1"
    }
  }
}
```

Add a `workspace_requirement` resource into your Terraform configuration to use this provider for guarding against applies in the wrong workspace.

```hcl
resource "workspace_requirement" "x" {}
```

Optionally update the `not_workspace` provider configuration:

```hcl
provider "workspace" {
  # Disable plans for the "default" workspace.
  not_workspace = "default"
}
```

If using the provider without an explicit provider block you may notice [this warning](https://github.com/hashicorp/terraform/issues/20121#issuecomment-457856988) attached to the expected "workspace restricted" error or when the workspace name matches the "not_workspace" argument.

> Provider "erran/workspace" requires explicit configuration.
>
> Add a provider block to the root module and configure the provider's required arguments as described in the provider documentation.

This can be suppressed by adding `provider "workspace" {}` block.
