package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新终端设备孪生参数配置
type TwinUpdateDetail struct {
	Twin *ValueInTwin `json:"twin,omitempty"`

	PropertyVisitors *TwinUpdateDetailPropertyVisitors `json:"property_visitors,omitempty"`
}

func (o TwinUpdateDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TwinUpdateDetail struct{}"
	}

	return strings.Join([]string{"TwinUpdateDetail", string(data)}, " ")
}
