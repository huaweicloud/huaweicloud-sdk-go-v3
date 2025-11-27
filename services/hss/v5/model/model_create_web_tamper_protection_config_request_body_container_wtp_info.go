package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWebTamperProtectionConfigRequestBodyContainerWtpInfo **参数解释**: 容器网页防篡改相关配置信息 **约束限制**: type值为“container_wtp”容器网页防篡改时有效 **取值范围**: 不涉及 **默认取值**: 不涉及
type CreateWebTamperProtectionConfigRequestBodyContainerWtpInfo struct {

	// **参数解释**: 防护类型 **约束限制**: 不涉及 **取值范围**: - cluster：对集群下的容器进行网页防篡改防护 - host：对主机下的容器进行网页防篡改防护  **默认取值**: 不涉及
	ProtectType string `json:"protect_type"`

	// **参数解释**: 网站应用的名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	WebAppName string `json:"web_app_name"`

	// **参数解释**: 集群id **约束限制**: protect_type值为“cluster”时有效 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 主机id **约束限制**: protect_type值为“host”时有效 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 集群标签列表 **约束限制**: protect_type值为“cluster”时生效 **取值范围**: 最少0条，最多10条 **默认取值**: 不涉及
	ClusterLabelList *[]ClusterLabelInfo `json:"cluster_label_list,omitempty"`

	// **参数解释**: 防护的镜像列表 **约束限制**: 不涉及 **取值范围**: 最少1条，最多10条 **默认取值**: 不涉及
	ProtectImageList []ProtectImageInfo `json:"protect_image_list"`
}

func (o CreateWebTamperProtectionConfigRequestBodyContainerWtpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWebTamperProtectionConfigRequestBodyContainerWtpInfo struct{}"
	}

	return strings.Join([]string{"CreateWebTamperProtectionConfigRequestBodyContainerWtpInfo", string(data)}, " ")
}
