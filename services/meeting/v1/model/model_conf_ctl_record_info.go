package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会控操作信息
type ConfCtlRecordInfo struct {

	// 操作时间（UTC时间，单位毫秒）。
	OperateTime *int64 `json:"operateTime,omitempty" xml:"operateTime"`

	// 操作来源。
	OperateSource *string `json:"operateSource,omitempty" xml:"operateSource"`

	// 操作者。
	Operator *string `json:"operator,omitempty" xml:"operator"`

	// 操作描述。
	OperateCode *string `json:"operateCode,omitempty" xml:"operateCode"`

	// 被操作对象。
	OperationObject *string `json:"operationObject,omitempty" xml:"operationObject"`

	// 操作结果。
	OperateResult *string `json:"operateResult,omitempty" xml:"operateResult"`

	// 详情。
	Detail *string `json:"detail,omitempty" xml:"detail"`
}

func (o ConfCtlRecordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfCtlRecordInfo struct{}"
	}

	return strings.Join([]string{"ConfCtlRecordInfo", string(data)}, " ")
}
