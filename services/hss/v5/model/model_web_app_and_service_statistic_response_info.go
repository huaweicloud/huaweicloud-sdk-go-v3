package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebAppAndServiceStatisticResponseInfo struct {

	// WebAppAndService资产名称
	Name *string `json:"name,omitempty"`

	// WebAppAndService资产--具备该资产的主机数量
	Num *int32 `json:"num,omitempty"`
}

func (o WebAppAndServiceStatisticResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebAppAndServiceStatisticResponseInfo struct{}"
	}

	return strings.Join([]string{"WebAppAndServiceStatisticResponseInfo", string(data)}, " ")
}
