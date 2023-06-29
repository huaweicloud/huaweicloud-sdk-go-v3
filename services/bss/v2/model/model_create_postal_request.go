package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePostalRequest Request Object
type CreatePostalRequest struct {

	// 语言。中文：zh_CN 英文：en_US，缺省为zh_CN
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
