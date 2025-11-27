package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWebTamperProtectionConfigRequestInfoContainerWtpInfo **参数解释**: 容器网页防篡改相关配置信息 **约束限制**: type值为“container_wtp”容器网页防篡改时有效 **取值范围**: 不涉及 **默认取值**: 不涉及
type UpdateWebTamperProtectionConfigRequestInfoContainerWtpInfo struct {

	// **参数解释**: 集群标签列表 **约束限制**: protect_type值为“cluster”时生效 **取值范围**: 最少0条，最多10条 **默认取值**: 不涉及
	ClusterLabelList *[]ClusterLabelInfo `json:"cluster_label_list,omitempty"`
}

func (o UpdateWebTamperProtectionConfigRequestInfoContainerWtpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebTamperProtectionConfigRequestInfoContainerWtpInfo struct{}"
	}

	return strings.Join([]string{"UpdateWebTamperProtectionConfigRequestInfoContainerWtpInfo", string(data)}, " ")
}
