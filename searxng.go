package searxng

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Client struct {
	apiURL     *url.URL
	httpClient *http.Client
}

func wrapErrorf(format string, args ...any) error {
	return fmt.Errorf("searxng: "+format, args...)
}

type ClientOption struct {
	// URL is required, The URL of searxng API.
	URL string
	// HTTPClient is optional, The HTTP client to use.
	HTTPClient *http.Client
}

func NewClient(opt *ClientOption) (*Client, error) {
	u, err := url.Parse(opt.URL)
	if err != nil {
		return nil, wrapErrorf("bad url: %w", err)
	}
	httpClient := http.DefaultClient
	if opt.HTTPClient != nil {
		httpClient = opt.HTTPClient
	}
	return &Client{
		u,
		httpClient,
	}, nil
}

// SearchInput is the input for the search function.
// ref: https://docs.searxng.org/dev/search_api.html
type SearchInput struct {
	// Query is required. The search query.
	Query string
	// Categories is optional. Specifies the active search categories.
	Categories []Category
	// Engines is optional. Specifies the active search engines.
	Engines []SearchEngine
	// Language is optional. Code of the language to use for the search.
	Language *string
	// SafeSearch is optional. Search page number.
	PageNumber *uint
	// TimeRange is optional. Time range of the search if the search engine supports it.
	TimeRange *TimeRange
	// SafeSearch is optional. Level of search result filter.
	SafeSearch *SafeSearchLevel
}

func (c *Client) Search(ctx context.Context, input *SearchInput) (*SearchOutput, error) {
	requestURL := *c.apiURL
	params := url.Values{}
	params.Set("q", input.Query)
	params.Set("format", "json")
	if input.Categories != nil {
		params.Set("categories", joinComma(input.Categories))
	}
	if input.Engines != nil {
		params.Set("engines", joinComma(input.Engines))
	}
	if input.Language != nil {
		params.Set("language", *input.Language)
	}
	if input.PageNumber != nil {
		params.Set("pageno", strconv.FormatUint(uint64(*input.PageNumber), 10))
	}
	if input.TimeRange != nil {
		params.Set("time_range", string(*input.TimeRange))
	}
	if input.SafeSearch != nil {
		params.Set("safesearch", strconv.FormatInt(int64(*input.SafeSearch), 10))
	}
	requestURL.RawQuery = params.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL.String(), nil)
	if err != nil {
		return nil, wrapErrorf("creating search request: %w", err)
	}
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, wrapErrorf("http error: %w", err)
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		return nil, wrapErrorf("bad http status: %d", httpResp.StatusCode)
	}

	var output SearchOutput
	if err := json.NewDecoder(httpResp.Body).Decode(&output); err != nil {
		return nil, wrapErrorf("decoding response: %w", err)
	}

	return &output, nil
}

func joinComma[T ~string](s []T) string {
	var b strings.Builder
	for i, v := range s {
		if i != 0 {
			b.WriteString(",")
		}
		b.WriteString(string(v))
	}
	return b.String()
}
