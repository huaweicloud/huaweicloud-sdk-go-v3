package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddSshKeyRequestBody struct {

	// 密钥
	Key string `json:"key" xml:"key"`

	// 密钥名称
	Title string `json:"title" xml:"title"`
}

func (o AddSshKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSshKeyRequestBody struct{}"
	}

	return strings.Join([]string{"AddSshKeyRequestBody", string(data)}, " ")
}
