package model

import (
	"encoding/json"

	"strings"
)

// 请求消息的数据部分。
type IvsExtentionByIdCardImageRequestBodyData struct {
	// 请求列表，用于支持批量调用。目前暂时只支持单个数据查询。

	ReqData *[]ExtentionReqDataByIdCardImage `json:"req_data,omitempty"`
}

func (o IvsExtentionByIdCardImageRequestBodyData) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IvsExtentionByIdCardImageRequestBodyData struct{}"
	}

	return strings.Join([]string{"IvsExtentionByIdCardImageRequestBodyData", string(data)}, " ")
}
