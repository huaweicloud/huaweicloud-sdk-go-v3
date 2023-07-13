package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IvsStandardByVideoAndNameAndIdRequestBodyData 请求消息的数据部分。
type IvsStandardByVideoAndNameAndIdRequestBodyData struct {

	// 请求列表，用于支持批量调用。目前暂时只支持单个数据查询。
	ReqData *[]StandardReqDataByVideoAndNameAndId `json:"req_data,omitempty"`
}

func (o IvsStandardByVideoAndNameAndIdRequestBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsStandardByVideoAndNameAndIdRequestBodyData struct{}"
	}

	return strings.Join([]string{"IvsStandardByVideoAndNameAndIdRequestBodyData", string(data)}, " ")
}
