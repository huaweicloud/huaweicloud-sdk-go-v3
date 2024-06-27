package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateChInstanceInfoTagsInfoTags struct {

	// 标签键。
	Key *string `json:"key,omitempty"`

	// 标签值。
	Value *string `json:"value,omitempty"`
}

func (o CreateChInstanceInfoTagsInfoTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChInstanceInfoTagsInfoTags struct{}"
	}

	return strings.Join([]string{"CreateChInstanceInfoTagsInfoTags", string(data)}, " ")
}
