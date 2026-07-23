package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterTag struct {

	// 标签键
	Value *string `json:"value,omitempty"`

	// 标签值
	Key string `json:"key"`
}

func (o ClusterTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterTag struct{}"
	}

	return strings.Join([]string{"ClusterTag", string(data)}, " ")
}
