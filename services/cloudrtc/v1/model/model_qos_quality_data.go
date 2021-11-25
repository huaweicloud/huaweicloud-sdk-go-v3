package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QosQualityData struct {
	// 用户id

	Uid *string `json:"uid,omitempty"`
	// 对端的用户ID，为0时表示本端上行数据

	Peerid *string `json:"peerid,omitempty"`
	// 指标ID

	Mid *string `json:"mid,omitempty"`
	// 时间戳及相应时间的指标数值

	Data *[]TimeFloatValueData `json:"data,omitempty"`
}

func (o QosQualityData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QosQualityData struct{}"
	}

	return strings.Join([]string{"QosQualityData", string(data)}, " ")
}
