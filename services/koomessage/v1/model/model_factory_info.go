package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FactoryInfo 各终端厂商的审核状态。
type FactoryInfo struct {

	// 厂商类型。  - HUAWEI：表示华为厂商 - XIAOMI：表示小米厂商 - OPPO：表示OPPO厂商 - VIVO：表示VIVO厂商 - MEIZU：表示魅族厂商 - SAMSUNG：表示三星厂商
	FactoryType string `json:"factory_type"`

	// 模板状态。  - 1：激活  - 其他：未激活
	State int32 `json:"state"`

	// 厂商版本
	Version *string `json:"version,omitempty"`

	// 智能信息模板ID
	TplId *string `json:"tpl_id,omitempty"`
}

func (o FactoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FactoryInfo struct{}"
	}

	return strings.Join([]string{"FactoryInfo", string(data)}, " ")
}
