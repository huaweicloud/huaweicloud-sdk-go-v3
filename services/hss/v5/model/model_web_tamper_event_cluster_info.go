package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperEventClusterInfo **参数解释**: 网页防篡改事件信息对应的集群信息 **取值范围**: 不涉及
type WebTamperEventClusterInfo struct {

	// **参数解释**： 集群ID **取值范围**： 字符长度1-128位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 集群名称 **取值范围**： 字符长度1-256位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 集群类型 **取值范围**： 字符长度1-128位
	ClusterType *string `json:"cluster_type,omitempty"`
}

func (o WebTamperEventClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperEventClusterInfo struct{}"
	}

	return strings.Join([]string{"WebTamperEventClusterInfo", string(data)}, " ")
}
