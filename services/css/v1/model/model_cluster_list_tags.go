package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群标签。
type ClusterListTags struct {

	// 集群标签的key值。
	Key *string `json:"key,omitempty" xml:"key"`

	// 集群标签的value值。
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o ClusterListTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterListTags struct{}"
	}

	return strings.Join([]string{"ClusterListTags", string(data)}, " ")
}
