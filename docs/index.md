# Provider: workspace
The "workspace" provider provides the ability to add restrictions of what workspaces to allow plan/applies in.

The motivation for this project are two recurring use cases I've seen used in different projects being managed with Terraform are:

1. Deployment environment workspaces: `qa`, `staging`, and `production`.
2. Regional workspaces: `us-east-1` and `eu-central-1`.

You can see that in either case you wouldn't have a "default" workspace.

You could just pick one environment and call it "default" but now you need to document that decision.

Now, what happens when you accidentally apply in the "default" workspace? Duplicate resources? Having to cherry-pick what to delete manually because you're not brave enough to use `terraform destroy` to delete dozens to hundreds of duplicate resources?

This provider allows you to:

1. Prevent Terraform from planning or applying any resources in the incorrect workspace. (provider level)
<!--2. Prevent Terraform from applying specific resources by leveraging `workspace_requirement` with Terraform's own `dependencies = [...]` resource argument. -->

-> The provider level settings will be verified before *any* resources have been planned. This means you can prevent plan/applies outside expected workspaces.

!> *At least one resource* is required to prevent Terraform ignoring the provider block.

You can use the provider-level `not_workspace` setting:
```hcl
provider "workspace" {
  not_workspace = "default"
}

# NOTE: At least one resource must be declared.
resource "workspace_requirement" "not_default" {}
```

See the [workspace_requirement](/docs/resources/requirement) resource for more details.
