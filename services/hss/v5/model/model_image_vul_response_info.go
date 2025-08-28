package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageVulResponseInfo **参数解释**: 镜像漏洞列表
type ImageVulResponseInfo struct {

	// **参数解释**: 镜像id **取值范围**: 字符长度0-512
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-512
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 漏洞id **取值范围**: 字符长度0-512
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞名称 **取值范围**: 字符长度0-512
	VulName *string `json:"vul_name,omitempty"`
}

func (o ImageVulResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageVulResponseInfo struct{}"
	}

	return strings.Join([]string{"ImageVulResponseInfo", string(data)}, " ")
}
