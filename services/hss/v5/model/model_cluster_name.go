package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterName 集群名称
type ClusterName struct {
}

func (o ClusterName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterName struct{}"
	}

	return strings.Join([]string{"ClusterName", string(data)}, " ")
}
