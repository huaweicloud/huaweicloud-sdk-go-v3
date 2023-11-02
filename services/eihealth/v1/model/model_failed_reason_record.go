package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailedReasonRecord struct {

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	// 失败数量
	Count *int32 `json:"count,omitempty"`
}

func (o FailedReasonRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedReasonRecord struct{}"
	}

	return strings.Join([]string{"FailedReasonRecord", string(data)}, " ")
}
