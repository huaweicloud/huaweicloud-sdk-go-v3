package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IvsStandardByNameAndIdResponseBodyResult 调用返回结果。
type IvsStandardByNameAndIdResponseBodyResult struct {

	// 子服务名称。
	ServiceName string `json:"service_name"`

	// 成功的结果数量，与resp_data字段对应。
	Count int32 `json:"count"`

	// 请求列表，用于支持批量调用。目前暂时只支持单个数据查询。
	RespData []StandardRespDataByNameAndId `json:"resp_data"`
}

func (o IvsStandardByNameAndIdResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsStandardByNameAndIdResponseBodyResult struct{}"
	}

	return strings.Join([]string{"IvsStandardByNameAndIdResponseBodyResult", string(data)}, " ")
}
