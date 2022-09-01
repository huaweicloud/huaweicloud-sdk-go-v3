package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddTagsResponse struct {

	// 标签名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 备注
	Message *string `json:"message,omitempty" xml:"message"`

	Commit *CommitRepoV2 `json:"commit,omitempty" xml:"commit"`
}

func (o AddTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTagsResponse struct{}"
	}

	return strings.Join([]string{"AddTagsResponse", string(data)}, " ")
}
