package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateClusterSpec struct {

	// 集群所在国家信息
	Country *string `json:"country,omitempty"`

	// 集群所在城市信息
	City *string `json:"city,omitempty"`
}

func (o UpdateClusterSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterSpec struct{}"
	}

	return strings.Join([]string{"UpdateClusterSpec", string(data)}, " ")
}
