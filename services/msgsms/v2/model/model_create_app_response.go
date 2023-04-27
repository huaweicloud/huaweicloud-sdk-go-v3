package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAppResponse struct {

	// 应用KEY
	AppKey *string `json:"app_key,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 应用主键ID
	Id *string `json:"id,omitempty"`

	// Application Secret，应用密钥
	AppSecret      *string `json:"app_secret,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppResponse struct{}"
	}

	return strings.Join([]string{"CreateAppResponse", string(data)}, " ")
}
