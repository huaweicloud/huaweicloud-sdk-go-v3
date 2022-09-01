package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddDeployKeyRequestBody struct {

	// 部署key的来源
	Application string `json:"application" xml:"application"`

	// 部署key是否可以推送代码
	CanPush bool `json:"can_push" xml:"can_push"`

	// 部署key
	Key string `json:"key" xml:"key"`

	// 部署key名称
	KeyTitle string `json:"key_title" xml:"key_title"`
}

func (o AddDeployKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeployKeyRequestBody struct{}"
	}

	return strings.Join([]string{"AddDeployKeyRequestBody", string(data)}, " ")
}
