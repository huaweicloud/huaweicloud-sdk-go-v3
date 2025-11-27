package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperProtectionConfigResponseInfoContainerWtpInfo **参数解释**: 容器网页防篡改信息 **取值范围**: 不涉及
type WebTamperProtectionConfigResponseInfoContainerWtpInfo struct {

	// **参数解释**： 网站应用名称 **取值范围**： 字符长度1-128位
	WebAppName *string `json:"web_app_name,omitempty"`

	// **参数解释**： 镜像名称 **取值范围**： 字符长度1-512位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**： 镜像版本 **取值范围**： 字符长度1-64位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**： 镜像类型 **取值范围**： - registry 仓库镜像 - local 本地镜像 - custom 自定义镜像
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**: 防护类型 **取值范围**: - cluster：对集群下的容器进行网页防篡改防护 - host：对主机下的容器进行网页防篡改防护
	ProtectType *string `json:"protect_type,omitempty"`

	ClusterInfo *WebTamperProtectionConfigResponseInfoContainerWtpInfoClusterInfo `json:"cluster_info,omitempty"`

	// **参数解释**: 集群标签列表 **取值范围**: 最少0条，最多10条
	ClusterLabelList *[]WebTamperProtectionConfigResponseInfoContainerWtpInfoClusterLabelList `json:"cluster_label_list,omitempty"`

	HostInfo *WebTamperProtectionConfigResponseInfoContainerWtpInfoHostInfo `json:"host_info,omitempty"`
}

func (o WebTamperProtectionConfigResponseInfoContainerWtpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperProtectionConfigResponseInfoContainerWtpInfo struct{}"
	}

	return strings.Join([]string{"WebTamperProtectionConfigResponseInfoContainerWtpInfo", string(data)}, " ")
}
