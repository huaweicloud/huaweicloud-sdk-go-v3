package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddTagsRequest struct {

	// 标签名称
	TagName string `json:"tag_name" xml:"tag_name"`

	// 分支名称
	Ref string `json:"ref" xml:"ref"`

	// 备注
	Message *string `json:"message,omitempty" xml:"message"`
}

func (o AddTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTagsRequest struct{}"
	}

	return strings.Join([]string{"AddTagsRequest", string(data)}, " ")
}
