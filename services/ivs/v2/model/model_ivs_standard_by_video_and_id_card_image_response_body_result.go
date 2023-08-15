package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IvsStandardByVideoAndIdCardImageResponseBodyResult 调用返回结果。
type IvsStandardByVideoAndIdCardImageResponseBodyResult struct {

	// 子服务名称。
	ServiceName string `json:"service_name"`

	// 成功的结果数量，与resp_data字段对应。
	Count int32 `json:"count"`

	// 请求列表，用于支持批量调用。目前暂时只支持单个数据查询。
	RespData []RespDataByVideoAndIdCardImage `json:"resp_data"`
}

func (o IvsStandardByVideoAndIdCardImageResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsStandardByVideoAndIdCardImageResponseBodyResult struct{}"
	}

	return strings.Join([]string{"IvsStandardByVideoAndIdCardImageResponseBodyResult", string(data)}, " ")
}
