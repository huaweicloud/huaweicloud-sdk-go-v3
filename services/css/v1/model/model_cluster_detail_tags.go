package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterDetailTags 集群标签。
type ClusterDetailTags struct {

	// 集群标签的key值。
	Key *string `json:"key,omitempty"`

	// 集群标签的value值。
	Value *string `json:"value,omitempty"`
}

func (o ClusterDetailTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterDetailTags struct{}"
	}

	return strings.Join([]string{"ClusterDetailTags", string(data)}, " ")
}
