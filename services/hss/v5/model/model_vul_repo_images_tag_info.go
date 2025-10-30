package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulRepoImagesTagInfo 漏洞关联的镜像版本
type VulRepoImagesTagInfo struct {

	// **参数解释**: 版本id **取值范围**: 字符长度0-65535位
	TagId *string `json:"tag_id,omitempty"`

	// **参数解释**: 版本名称 **取值范围**: 字符长度0-65535位
	TagName *string `json:"tag_name,omitempty"`

	// **参数解释**: \"版本大小\" **取值范围**: 最小值0，最大值2147483547
	Size *int32 `json:"size,omitempty"`

	// **参数解释**: 版本app名称 **取值范围**: 字符长度0-65535位
	AppName *string `json:"app_name,omitempty"`
}

func (o VulRepoImagesTagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulRepoImagesTagInfo struct{}"
	}

	return strings.Join([]string{"VulRepoImagesTagInfo", string(data)}, " ")
}
