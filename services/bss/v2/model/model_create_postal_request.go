package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePostalRequest struct {
	// |参数名称：语言| |参数的约束及描述：中文：zh_CN 英文：en_US缺省为zh_CN|

	XLanguage *string `json:"X-Language,omitempty"`

	Body *AddPostalReq `json:"body,omitempty"`
}

func (o CreatePostalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostalRequest struct{}"
	}

	return strings.Join([]string{"CreatePostalRequest", string(data)}, " ")
}
