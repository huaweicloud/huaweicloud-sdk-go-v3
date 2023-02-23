package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValidityInterval struct {

	// 日
	Days *int32 `json:"days,omitempty"`

	// 时
	Hours *int32 `json:"hours,omitempty"`

	// 分
	Minutes *int32 `json:"minutes,omitempty"`
}

func (o ValidityInterval) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidityInterval struct{}"
	}

	return strings.Join([]string{"ValidityInterval", string(data)}, " ")
}
