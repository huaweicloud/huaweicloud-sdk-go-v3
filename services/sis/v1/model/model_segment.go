package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Segment struct {

	// 一句的起始时间戳，单位ms。
	StartTime int32 `json:"start_time" xml:"start_time"`

	// 一句的结束时间戳，单位ms。
	EndTime int32 `json:"end_time" xml:"end_time"`

	Result *TranscriberResult `json:"result" xml:"result"`
}

func (o Segment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Segment struct{}"
	}

	return strings.Join([]string{"Segment", string(data)}, " ")
}
