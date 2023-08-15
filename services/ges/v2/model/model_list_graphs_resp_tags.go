package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListGraphsRespTags struct {

	// 标签key
	Key *string `json:"key,omitempty"`

	// 标签value
	Value *string `json:"value,omitempty"`
}

func (o ListGraphsRespTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGraphsRespTags struct{}"
	}

	return strings.Join([]string{"ListGraphsRespTags", string(data)}, " ")
}
