package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowInstanceExtendProductInfoRespHourly struct {

	// 消息引擎的名称，该字段显示为kafka。
	Name *string `json:"name,omitempty"`

	// 消息引擎的版本，当前仅支持1.1.0和2.3.0。
	Version *string `json:"version,omitempty"`

	// 产品规格列表。
	Values *[]ShowInstanceExtendProductInfoRespValues `json:"values,omitempty"`
}

func (o ShowInstanceExtendProductInfoRespHourly) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceExtendProductInfoRespHourly struct{}"
	}

	return strings.Join([]string{"ShowInstanceExtendProductInfoRespHourly", string(data)}, " ")
}
