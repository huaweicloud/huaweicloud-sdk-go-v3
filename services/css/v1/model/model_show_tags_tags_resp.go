package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群标签。
type ShowTagsTagsResp struct {

	// 集群标签的key值。
	Key *string `json:"key,omitempty" xml:"key"`

	// 集群标签的value值。
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o ShowTagsTagsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagsTagsResp struct{}"
	}

	return strings.Join([]string{"ShowTagsTagsResp", string(data)}, " ")
}
