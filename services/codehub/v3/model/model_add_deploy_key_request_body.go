package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddDeployKeyRequestBody struct {

	// 部署使用的SSH密钥的来源
	Application string `json:"application"`

	// 部署使用的SSH密钥是否可以推送代码
	CanPush bool `json:"can_push"`

	// 部署使用的SSH密钥
	Key string `json:"key"`

	// 部署使用的SSH密钥名称
	KeyTitle string `json:"key_title"`
}

func (o AddDeployKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeployKeyRequestBody struct{}"
	}

	return strings.Join([]string{"AddDeployKeyRequestBody", string(data)}, " ")
}
