package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"golang.org/x/oauth2"
)

const COMMIT_MESSAGE = "Automated commit via custom Terraform provider."

func resourceRepoContent() *schema.Resource {
	return &schema.Resource{
		Create: resourceRepoContentCreate,
		Read:   resourceRepoContentRead,
		// Update: resourceRepoContentUpdate,
		Delete: resourceRepoContentDelete,

		Schema: map[string]*schema.Schema{
			"repo": {
				Type:     schema.TypeString,
				Required: true,
			},
			"file_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"file_content": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRepoContentCreate(d *schema.ResourceData, m interface{}) error {
	TOKEN := d.Get("token").(string)
	org := d.Get("organization").(string)
	repo := d.Get("repo").(string)
	file_path := d.Get("file_path").(string)
	file_content := d.Get("file_content").(string)

	tc, err := CreateOauth2Client(TOKEN)
	if err != nil {
		return err
	}

	client := github.NewClient(tc)
	err = UpdateFile(client, org, repo, file_path, file_content)
	if err != nil {
		return err
	}

	d.SetId(org + "/" + repo + "/" + file_path)
	return resourceRepoContentRead(d, m)
}

func resourceRepoContentRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

/*
func resourceRepoContentUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceRepoContentRead(d, m)
}
*/

func resourceRepoContentDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

// Authenticate and create OAuth2 client
func CreateOauth2Client(token string) (*http.Client, error) {
	if len(token) == 0 {
		return nil, errors.New("GitHub access token not defined.")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(context.Background(), ts)

	return tc, nil
}

// Retrieve file, append configuration, and prepare new commit
func UpdateFile(client *github.Client, org string, repo string, file_path string, file_content string) error {
	// Retrieve configuration file contents
	fileContent, _, _, err := client.Repositories.GetContents(context.Background(), org, repo, file_path, nil)
	if err != nil {
		return err
	}

	// Re-encode file with appended configuration
	encodedContent := []byte(file_content)
	commitMessage := COMMIT_MESSAGE

	// Apply new commit with file updates
	_, _, err = client.Repositories.UpdateFile(context.Background(), org, repo, file_path, &github.RepositoryContentFileOptions{
		Message: &commitMessage,
		Content: encodedContent,
		SHA:     fileContent.SHA,
	})
	if err != nil {
		return err
	}

	return nil
}
