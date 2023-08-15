package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IvsStandardByIdCardImageRequestBodyData 请求消息的数据部分。
type IvsStandardByIdCardImageRequestBodyData struct {

	// 请求列表，用于支持批量调用。目前暂时只支持单个数据查询。
	ReqData *[]ReqDataByIdCardImage `json:"req_data,omitempty"`
}

func (o IvsStandardByIdCardImageRequestBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsStandardByIdCardImageRequestBodyData struct{}"
	}

	return strings.Join([]string{"IvsStandardByIdCardImageRequestBodyData", string(data)}, " ")
}
