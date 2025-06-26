package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Credential struct {

	// 认证方式，目前只支持basic
	Type string `json:"type"`

	// 访问ID
	AccessKey string `json:"access_key"`

	// 访问密码
	AccessSecret string `json:"access_secret"`
}

func (o Credential) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Credential struct{}"
	}

	return strings.Join([]string{"Credential", string(data)}, " ")
}
