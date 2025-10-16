package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OtaModuleInfo OTA模块信息
type OtaModuleInfo struct {

	// OTA模块ID
	ModuleId *string `json:"module_id,omitempty"`

	// 资源空间ID
	AppId *string `json:"app_id,omitempty"`

	// OTA模块关联的产品ID
	ProductId *string `json:"product_id,omitempty"`

	// OTA模块关联的产品名称
	ProductName *string `json:"product_name,omitempty"`

	// OTA模块名称。
	ModuleName *string `json:"module_name,omitempty"`

	// OTA模块别名
	AliasName *string `json:"alias_name,omitempty"`

	// 用于描述模块的功能等信息
	Description *string `json:"description,omitempty"`

	// 创建OTA模块的时间，格式：\"yyyyMMdd'T'HHmmss'Z'\"。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o OtaModuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OtaModuleInfo struct{}"
	}

	return strings.Join([]string{"OtaModuleInfo", string(data)}, " ")
}
