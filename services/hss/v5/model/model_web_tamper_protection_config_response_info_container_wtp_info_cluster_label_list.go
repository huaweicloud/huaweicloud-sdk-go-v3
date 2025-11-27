package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperProtectionConfigResponseInfoContainerWtpInfoClusterLabelList **参数解释**: 集群资源标签 **取值范围**: 不涉及
type WebTamperProtectionConfigResponseInfoContainerWtpInfoClusterLabelList struct {

	// **参数解释**: 集群资源标签的键名 **取值范围**: 字符长度1-64位
	Key *string `json:"key,omitempty"`

	// **参数解释**: 集群资源标签的值 **取值范围**: 字符长度0-128位
	Value *string `json:"value,omitempty"`
}

func (o WebTamperProtectionConfigResponseInfoContainerWtpInfoClusterLabelList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperProtectionConfigResponseInfoContainerWtpInfoClusterLabelList struct{}"
	}

	return strings.Join([]string{"WebTamperProtectionConfigResponseInfoContainerWtpInfoClusterLabelList", string(data)}, " ")
}
