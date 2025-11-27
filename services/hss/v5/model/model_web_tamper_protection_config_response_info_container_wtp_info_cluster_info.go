package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperProtectionConfigResponseInfoContainerWtpInfoClusterInfo **参数解释**： 防护配置对应的集群信息 **取值范围**： 不涉及
type WebTamperProtectionConfigResponseInfoContainerWtpInfoClusterInfo struct {

	// **参数解释**： 集群ID **取值范围**： 字符长度0-128位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 集群名称 **取值范围**： 字符长度0-256位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 集群版本 **取值范围**： 字符长度0-256位
	ClusterVersion *string `json:"cluster_version,omitempty"`

	// **参数解释**： 集群类型 **取值范围**： - cce：cce集群 - ali：阿里云集群 - tencent：腾讯云集群 - azure：微软云集群 - aws：亚马逊集群 - self_built_hw：华为云自建集群 - self_built_idc：IDC自建集群
	ClusterType *string `json:"cluster_type,omitempty"`
}

func (o WebTamperProtectionConfigResponseInfoContainerWtpInfoClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperProtectionConfigResponseInfoContainerWtpInfoClusterInfo struct{}"
	}

	return strings.Join([]string{"WebTamperProtectionConfigResponseInfoContainerWtpInfoClusterInfo", string(data)}, " ")
}
