package confluence

import (
	"fmt"

	goconfluence "github.com/virtomize/confluence-go-api"
)

type ConfluenceClient struct {
	BaseURL   string
	APIToken  string
	Username  string
	apiClient *goconfluence.API
}

func NewConfluenceClient(baseURL, username, apiToken string) (*ConfluenceClient, error) {
	api, err := goconfluence.NewAPI(baseURL, username, apiToken)
	if err != nil {
		return nil, err
	}
	return &ConfluenceClient{
		BaseURL:   baseURL,
		APIToken:  apiToken,
		Username:  username,
		apiClient: api,
	}, nil
}

func (c *ConfluenceClient) CreatePage(page *CreateConfluencePage) (*goconfluence.Content, error) {
	content, err := c.apiClient.CreateContent(&goconfluence.Content{
		Type:  "page",
		Title: page.Title,
		Space: &goconfluence.Space{
			Key: page.SpaceKey,
		},
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          page.Body,
				Representation: "storage",
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (c *ConfluenceClient) UpdatePage(page *UpdateConfluencePage) (*goconfluence.Content, error) {
	content, err := c.apiClient.UpdateContent(&goconfluence.Content{
		ID:    page.PageID,
		Type:  "page",
		Title: page.Title,
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          page.Body,
				Representation: "storage",
			},
		},
		Version: &goconfluence.Version{
			Number: page.Version,
		},
	})
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (c *ConfluenceClient) GetPageByID(pageID string) (*goconfluence.Content, error) {
	content, err := c.apiClient.GetContentByID(pageID, goconfluence.ContentQuery{
		Expand: []string{"body.storage", "version"},
	})
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (c *ConfluenceClient) GetPageByTitle(spaceKey, title string) (*goconfluence.Content, error) {
	contents, err := c.apiClient.GetContent(goconfluence.ContentQuery{
		SpaceKey: spaceKey,
		Title:    title,
		Expand:   []string{"body.storage", "version"},
	})
	if err != nil {
		return nil, err
	}
	if len(contents.Results) == 0 {
		return nil, fmt.Errorf("page not found")
	}
	return &contents.Results[0], nil
}

func (c *ConfluenceClient) DeletePage(pageID string) error {
	_, err := c.apiClient.DelContent(pageID)
	return err
}

func (c *ConfluenceClient) GetSpacePages(params GetSpacePages) ([]goconfluence.Content, error) {
	contents, err := c.apiClient.GetContent(goconfluence.ContentQuery{
		SpaceKey: params.SpaceKey,
		Limit:    params.Limit,
		Start:    params.Start,
		Expand:   []string{"body.storage", "version"},
	})
	if err != nil {
		return nil, err
	}
	return contents.Results, nil
}
