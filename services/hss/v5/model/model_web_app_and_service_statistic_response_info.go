package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebAppAndServiceStatisticResponseInfo struct {

	// **参数解释**: web应用、web服务、数据库资产名称 **取值范围**: 字符长度0-256
	Name *string `json:"name,omitempty"`

	// **参数解释** web应用、web服务、数据库资产--具备该资产的主机数量 **取值范围** 最小值0，最大值300000
	Num *int32 `json:"num,omitempty"`
}

func (o WebAppAndServiceStatisticResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebAppAndServiceStatisticResponseInfo struct{}"
	}

	return strings.Join([]string{"WebAppAndServiceStatisticResponseInfo", string(data)}, " ")
}
