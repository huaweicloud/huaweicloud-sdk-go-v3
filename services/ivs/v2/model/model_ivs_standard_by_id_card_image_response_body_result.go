package model

import (
	"encoding/json"

	"strings"
)

// 调用返回结果。
type IvsStandardByIdCardImageResponseBodyResult struct {
	// 子服务名称。

	ServiceName *string `json:"service_name,omitempty"`
	// 成功的结果数量，与resp_data字段对应。

	Count *int32 `json:"count,omitempty"`
	// 请求列表，用于支持批量调用。目前暂时只支持单个数据查询。

	RespData *[]RespDataByIdCardImage `json:"resp_data,omitempty"`
}

func (o IvsStandardByIdCardImageResponseBodyResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IvsStandardByIdCardImageResponseBodyResult struct{}"
	}

	return strings.Join([]string{"IvsStandardByIdCardImageResponseBodyResult", string(data)}, " ")
}
