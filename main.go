package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

type WorkspaceConfig struct {
	NotWorkspace string
}

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: Provider,
	}
	plugin.Serve(&opts)
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"not_workspace": &schema.Schema{
				DefaultFunc: func() (interface{}, error) {
					return "default", nil
				},
				Description: "The workspace to ",
				Required:    true,
				Type:        schema.TypeString,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"workspace_requirement": resourceWorkspaceRequirement(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := WorkspaceConfig{
		NotWorkspace: d.Get("not_workspace").(string),
	}

	w, err := currentWorkspace()
	if err != nil {
		return &config, err
	} else if w == config.NotWorkspace {
		return &config, fmt.Errorf("[PROVIDER] Workspace was %v but expected %v", w, config.NotWorkspace)
	}
	return &config, nil
}

func resourceWorkspaceRequirementCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(fmt.Sprintf("%d", rand.Int()))
	return nil
}

func resourceWorkspaceRequirementRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceWorkspaceRequirementDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

func currentWorkspace() (string, error) {
	w := os.Getenv("TERRAFORM_WORKSPACE")
	if w != "" {
		return w, nil
	}

	if _, err := os.Stat(".terraform/environment"); err == nil {
		b, err := ioutil.ReadFile(".terraform/environment")
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	return "default", nil
}

func resourceWorkspaceRequirement() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkspaceRequirementCreate,
		Read:   resourceWorkspaceRequirementRead,
		Delete: resourceWorkspaceRequirementDelete,
		Schema: map[string]*schema.Schema{},
	}
}
