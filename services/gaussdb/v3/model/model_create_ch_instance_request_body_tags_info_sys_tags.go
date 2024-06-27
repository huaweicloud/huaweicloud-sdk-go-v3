package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateChInstanceRequestBodyTagsInfoSysTags struct {

	// 标签键。
	Key string `json:"key"`

	// 标签值。
	Value string `json:"value"`
}

func (o CreateChInstanceRequestBodyTagsInfoSysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChInstanceRequestBodyTagsInfoSysTags struct{}"
	}

	return strings.Join([]string{"CreateChInstanceRequestBodyTagsInfoSysTags", string(data)}, " ")
}
