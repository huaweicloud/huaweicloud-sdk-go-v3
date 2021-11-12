package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RespcodeBrokens struct {
	// checkPointFailed

	CheckPointFailed *[]int32 `json:"checkPointFailed,omitempty"`
	// error

	Error *[]int32 `json:"error,omitempty"`
	// othersFailed

	OthersFailed *[]int32 `json:"othersFailed,omitempty"`
	// parsedFailed

	ParsedFailed *[]int32 `json:"parsedFailed,omitempty"`
	// refusedFailed

	RefusedFailed *[]int32 `json:"refusedFailed,omitempty"`
	// success

	Success *[]int32 `json:"success,omitempty"`
	// timeout

	Timeout *[]int32 `json:"timeout,omitempty"`
}

func (o RespcodeBrokens) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespcodeBrokens struct{}"
	}

	return strings.Join([]string{"RespcodeBrokens", string(data)}, " ")
}
