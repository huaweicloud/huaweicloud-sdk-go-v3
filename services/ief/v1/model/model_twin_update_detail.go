package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TwinUpdateDetail 更新终端设备孪生参数配置
type TwinUpdateDetail struct {

	// 终端设备动态属性
	Twin map[string]ValueInTwin `json:"twin,omitempty"`

	// 孪生属性配置，与access_protocol关联。
	PropertyVisitors map[string]ValueInPropertyVisitors `json:"property_visitors,omitempty"`
}

func (o TwinUpdateDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TwinUpdateDetail struct{}"
	}

	return strings.Join([]string{"TwinUpdateDetail", string(data)}, " ")
}
