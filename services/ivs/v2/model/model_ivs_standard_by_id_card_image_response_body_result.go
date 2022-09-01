package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 调用返回结果。
type IvsStandardByIdCardImageResponseBodyResult struct {

	// 子服务名称。
	ServiceName *string `json:"service_name,omitempty" xml:"service_name"`

	// 成功的结果数量，与resp_data字段对应。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 请求列表，用于支持批量调用。目前暂时只支持单个数据查询。
	RespData *[]RespDataByIdCardImage `json:"resp_data,omitempty" xml:"resp_data"`
}

func (o IvsStandardByIdCardImageResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsStandardByIdCardImageResponseBodyResult struct{}"
	}

	return strings.Join([]string{"IvsStandardByIdCardImageResponseBodyResult", string(data)}, " ")
}
