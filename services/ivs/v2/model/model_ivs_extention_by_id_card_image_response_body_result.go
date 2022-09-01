package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 调用返回结果。
type IvsExtentionByIdCardImageResponseBodyResult struct {

	// 子服务名称。
	ServiceName *string `json:"service_name,omitempty" xml:"service_name"`

	// 成功的结果数量，与resp_data字段对应。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 请求列表，用于支持批量调用。目前暂时只支持单个数据查询。
	RespData *[]ExtentionRespDataByIdCardImage `json:"resp_data,omitempty" xml:"resp_data"`
}

func (o IvsExtentionByIdCardImageResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsExtentionByIdCardImageResponseBodyResult struct{}"
	}

	return strings.Join([]string{"IvsExtentionByIdCardImageResponseBodyResult", string(data)}, " ")
}
