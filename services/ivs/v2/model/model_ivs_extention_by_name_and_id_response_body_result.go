package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 调用返回结果。
type IvsExtentionByNameAndIdResponseBodyResult struct {
	// 子服务名称。

	ServiceName *string `json:"service_name,omitempty"`
	// 成功的结果数量，与resp_data字段对应。

	Count *int32 `json:"count,omitempty"`
	// 请求列表，用于支持批量调用。目前暂时只支持单个数据查询。

	RespData *[]ExtentionRespDataByNameAndId `json:"resp_data,omitempty"`
}

func (o IvsExtentionByNameAndIdResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsExtentionByNameAndIdResponseBodyResult struct{}"
	}

	return strings.Join([]string{"IvsExtentionByNameAndIdResponseBodyResult", string(data)}, " ")
}
