package fitbit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseUrl = "https://api.fitbit.com/1/"
	userAgent      = "go-fitbit"
)

type Client struct {
	client *http.Client

	BaseURL   *url.URL
	UserAgent string

	HeartRate *HeartRateService
	Sleep     *SleepService
	Activity  *ActivityService
	Body      *BodyService
	Food      *FoodService
	User      *UserService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseUrl, _ := url.Parse(defaultBaseUrl)

	c := &Client{client: httpClient, BaseURL: baseUrl, UserAgent: userAgent}

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil

}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = withContext(ctx, req)

	// rateLimitCategory := category(req.URL.Path)
	// // If we've hit rate limit, don't make further requests before Reset time.
	// if err := c.checkRateLimitBeforeDo(req, rateLimitCategory); err != nil {
	// 	return &Response{
	// 		Response: err.Response,
	// 		Rate:     err.Rate,
	// 	}, err
	// }

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = url.String()
				return nil, e
			}
		}

		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)

	// c.rateMu.Lock()
	// c.rateLimits[rateLimitCategory] = response.Rate
	// c.rateMu.Unlock()

	// err = CheckResponse(resp)
	// if err != nil {
	// 	// Even though there was an error, we still return the response
	// 	// in case the caller wants to inspect it further.
	// 	// However, if the error is AcceptedError, decode it below before
	// 	// returning from this function and closing the response body.
	// 	if _, ok := err.(*AcceptedError); !ok {
	// 		return response, err
	// 	}
	// }

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return response, err
}
