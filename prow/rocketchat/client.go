/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rocketchat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"

	"k8s.io/test-infra/prow/clientutil"
)

// Client allows you to provide connection to RocketChat API Server
// It contains a secret webhookURL that allows to authenticate connection to post and work with channels in the domain
type Client struct {
	// If logger is non-nil, log all method calls with it.
	logger clientutil.Logger

	fake bool

	webhookURL string
}

// NewClient creates a RocketChat client with a secret webhookURL.
func NewClient(webhook func() []byte) *Client {
	webhookURL := string(webhook())
	return &Client{
		logger:     logrus.WithField("client", "rocketchat"),
		webhookURL: webhookURL,
	}
}

// NewFakeClient returns a client that takes no actions.
func NewFakeClient() *Client {
	return &Client{
		fake: true,
	}
}

func (sl *Client) log(methodName string, args ...interface{}) {
	if sl.logger == nil {
		return
	}
	var as []string
	for _, arg := range args {
		as = append(as, fmt.Sprintf("%v", arg))
	}
	sl.logger.Debugf("%s(%s)", methodName, strings.Join(as, ", "))
}

func (sl *Client) postMessage(url string, uv *url.Values) error {
	resp, err := http.PostForm(url, *uv)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	apiResponse := struct {
		Ok    bool   `json:"ok"`
		Error string `json:"error"`
	}{}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return fmt.Errorf("API returned invalid JSON (%q): %w", string(body), err)
	}

	if resp.StatusCode != 200 || !apiResponse.Ok {
		return fmt.Errorf("request failed: %s", apiResponse.Error)
	}

	return nil
}

// WriteMessage adds text to channel
func (sl *Client) WriteMessage(text, channel string) error {
	sl.log("WriteMessage", text, channel)
	if sl.fake {
		return nil
	}

	uv := url.Values{}
	if channel != "" {
		uv.Add("channel", channel)
	}
	uv.Add("text", text)

	if err := sl.postMessage(sl.webhookURL, &uv); err != nil {
		return fmt.Errorf("failed to post message to %s: %w", channel, err)
	}
	return nil
}
