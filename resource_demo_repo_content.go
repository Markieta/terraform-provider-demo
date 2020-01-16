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
const REFS_PREFIX = "refs/heads/"

func resourceRepoContent() *schema.Resource {
	return &schema.Resource{
		Create: resourceRepoContentCreate,
		Read:   resourceRepoContentRead,
		Update: resourceRepoContentUpdate,
		Delete: resourceRepoContentDelete,

		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_TOKEN", nil),
				Description: "The OAuth token used to connect to GitHub.",
			},
			"organization": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_ORGANIZATION", nil),
				Description: "The GitHub organization name to manage.",
			},
			"repo": {
				Type:     schema.TypeString,
				Required: true,
			},
			"branch": {
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
	branch := d.Get("branch").(string)
	file_path := d.Get("file_path").(string)
	file_content := d.Get("file_content").(string)

	tc, err := CreateOauth2Client(TOKEN)
	if err != nil {
		return err
	}

	client := github.NewClient(tc)
	err = UpdateFile(client, org, repo, branch, file_path, file_content)
	if err != nil {
		return err
	}

	d.SetId(org + "/" + repo + "/" + branch + "/" + file_path)
	return resourceRepoContentRead(d, m)
}

func resourceRepoContentRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceRepoContentUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceRepoContentRead(d, m)
}

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

// Retrieve branch reference and write (commit) file
func UpdateFile(client *github.Client, org string, repo string, branch string, file_path string, file_content string) error {
	// Retrieve branch reference
	baseRef, _, err := client.Git.GetRef(context.Background(), org, repo, REFS_PREFIX+branch)
	if err != nil {
		return err
	}

	encodedContent := []byte(file_content)
	commitMessage := COMMIT_MESSAGE

	// Apply new commit with file updates
	_, _, err = client.Repositories.UpdateFile(context.Background(), org, repo, file_path, &github.RepositoryContentFileOptions{
		Message: &commitMessage,
		Content: encodedContent,
		SHA:     baseRef.Object.SHA,
	})
	if err != nil {
		return err
	}

	return nil
}
