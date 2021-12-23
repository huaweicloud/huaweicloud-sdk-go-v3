package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 请求消息的数据部分。
type IvsExtentionByNameAndIdRequestBodyData struct {
	// 请求列表，用于支持批量调用。目前暂时只支持单个数据查询。

	ReqData *[]ExtentionReqDataByNameAndId `json:"req_data,omitempty"`
}

func (o IvsExtentionByNameAndIdRequestBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsExtentionByNameAndIdRequestBodyData struct{}"
	}

	return strings.Join([]string{"IvsExtentionByNameAndIdRequestBodyData", string(data)}, " ")
}
