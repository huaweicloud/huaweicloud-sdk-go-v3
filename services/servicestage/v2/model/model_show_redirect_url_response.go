package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRedirectUrlResponse struct {

	// 授权重定向URL。
	Url            *string `json:"url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRedirectUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedirectUrlResponse struct{}"
	}

	return strings.Join([]string{"ShowRedirectUrlResponse", string(data)}, " ")
}
