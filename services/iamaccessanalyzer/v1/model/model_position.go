package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Position 策略中的位置。
type Position struct {

	// 位置的行号，从1开始。
	Line int32 `json:"line"`

	// 位置的列号，从0开始。
	Column int32 `json:"column"`

	// 策略中与位置对应的偏移量，从0开始。
	Offset int32 `json:"offset"`
}

func (o Position) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Position struct{}"
	}

	return strings.Join([]string{"Position", string(data)}, " ")
}
