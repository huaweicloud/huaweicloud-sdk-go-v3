package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会控操作信息
type ConfCtlRecordInfo struct {

	// 操作时间（UTC时间，单位毫秒）。
	OperateTime *int64 `json:"operateTime,omitempty"`

	// 操作来源。
	OperateSource *string `json:"operateSource,omitempty"`

	// 操作者。
	Operator *string `json:"operator,omitempty"`

	// 操作描述。
	OperateCode *string `json:"operateCode,omitempty"`

	// 被操作对象。
	OperationObject *string `json:"operationObject,omitempty"`

	// 操作结果。
	OperateResult *string `json:"operateResult,omitempty"`

	// 详情。
	Detail *string `json:"detail,omitempty"`
}

func (o ConfCtlRecordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfCtlRecordInfo struct{}"
	}

	return strings.Join([]string{"ConfCtlRecordInfo", string(data)}, " ")
}
