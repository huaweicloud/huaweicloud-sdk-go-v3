package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticSeverityV2 struct {

	// 致命问题数
	Critical *int32 `json:"critical,omitempty" xml:"critical"`

	// 严重问题数
	Major *int32 `json:"major,omitempty" xml:"major"`

	// 一般问题数
	Minor *int32 `json:"minor,omitempty" xml:"minor"`

	// 提示问题数
	Suggestion *int32 `json:"suggestion,omitempty" xml:"suggestion"`
}

func (o StatisticSeverityV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticSeverityV2 struct{}"
	}

	return strings.Join([]string{"StatisticSeverityV2", string(data)}, " ")
}
