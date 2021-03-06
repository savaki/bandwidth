package bandwidth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Error struct {
	StatusCode  int    `json:"-"`
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

func (e Error) Error() string {
	return e.Type + ": " + e.Description
}

type client struct {
	codebase string
	username string
	password string
}

func newClient(codebase, accountID, username, password string) *client {
	return &client{
		codebase: strings.TrimRight(codebase, "/") + "/" + accountID,
		username: username,
		password: password,
	}
}

func (c *client) Do(ctx context.Context, method, path string, body, v interface{}) error {
	var r io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal body for request, %v %v: %w", method, path, err)
		}
		r = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, c.codebase+path, r)
	if err != nil {
		return fmt.Errorf("failed to create request, %v %v: %w", method, path, err)
	}
	req = req.WithContext(ctx)
	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed api call, %v %v: %w", method, path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		io.Copy(os.Stdout, resp.Body)
		var e Error
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return fmt.Errorf("retrieved error for request, %v %v: %w", method, path, err)
		}
		e.StatusCode = resp.StatusCode
		return e
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("failed to decode response for %v %v: %w", method, path, err)
		}
	}

	return nil
}

func (c *client) Delete(ctx context.Context, path string, body, v interface{}) error {
	return c.Do(ctx, http.MethodDelete, path, body, v)
}

func (c *client) Get(ctx context.Context, path string, v interface{}) error {
	return c.Do(ctx, http.MethodGet, path, nil, v)
}

func (c *client) Post(ctx context.Context, path string, body, v interface{}) error {
	return c.Do(ctx, http.MethodPost, path, body, v)
}

func (c *client) Put(ctx context.Context, path string, body, v interface{}) error {
	return c.Do(ctx, http.MethodPut, path, body, v)
}
