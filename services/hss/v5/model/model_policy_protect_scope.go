package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyProtectScope 策略防护范围
type PolicyProtectScope struct {

	// 集群名称
	ClusterName string `json:"cluster_name"`

	// 集群id
	ClusterId string `json:"cluster_id"`

	// 防护镜像
	Images []string `json:"images"`

	// 防护namespace
	Namespaces []string `json:"namespaces"`

	// 防护labels
	Labels []string `json:"labels"`
}

func (o PolicyProtectScope) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyProtectScope struct{}"
	}

	return strings.Join([]string{"PolicyProtectScope", string(data)}, " ")
}
