package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiTaskResetBody struct {

	// 任务重置开始时间，UTC时间戳，允许为空
	DateFrom *int64 `json:"date_from,omitempty"`
}

func (o MultiTaskResetBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiTaskResetBody struct{}"
	}

	return strings.Join([]string{"MultiTaskResetBody", string(data)}, " ")
}
