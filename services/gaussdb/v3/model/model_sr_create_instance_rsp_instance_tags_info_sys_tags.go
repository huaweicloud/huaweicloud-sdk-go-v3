package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SrCreateInstanceRspInstanceTagsInfoSysTags struct {

	// 标签键。
	Key *string `json:"key,omitempty"`

	// 标签值。
	Value *string `json:"value,omitempty"`
}

func (o SrCreateInstanceRspInstanceTagsInfoSysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SrCreateInstanceRspInstanceTagsInfoSysTags struct{}"
	}

	return strings.Join([]string{"SrCreateInstanceRspInstanceTagsInfoSysTags", string(data)}, " ")
}
