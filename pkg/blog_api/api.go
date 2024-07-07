package blog_api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	APP_Key    = "CYY"
	APP_SECRET = "CYY"
)

type AccessToken struct {
	Token string `json:"token"`
}

func (a *API) name() {

}

type API struct {
	URL string
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

func (a *API) getAccessToken(ctx context.Context) (string, error) {
	url := fmt.Sprintf("%s?app_key=%s&app_secret=%s", "auth", APP_Key, APP_SECRET)
	body, err := a.httpGet(ctx, url)
	if err != nil {
		return "", err
	}

	var accessToken AccessToken
	_ = json.Unmarshal(body, &accessToken)
	return accessToken.Token, nil
}

func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	//accessToken, err := a.getAccessToken(ctx)
	//if err != nil {
	//	return nil, err
	//}

	//body, err := a.httpGet(ctx, fmt.Sprintf("%s?token=%s&name=%s", "api/tags", accessToken, name))
	body, err := a.httpGet(ctx, fmt.Sprintf("%s?name=%s", "api/tags", name))
	if err != nil {
		return nil, err
	}

	return body, nil
}
