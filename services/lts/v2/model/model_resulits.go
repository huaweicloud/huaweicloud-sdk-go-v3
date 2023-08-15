package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resulits struct {

	// 时间戳，毫秒时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 流量，byte
	Value *float64 `json:"value,omitempty"`
}

func (o Resulits) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resulits struct{}"
	}

	return strings.Join([]string{"Resulits", string(data)}, " ")
}
