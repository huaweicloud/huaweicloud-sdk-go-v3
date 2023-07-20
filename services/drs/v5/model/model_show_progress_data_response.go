package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProgressDataResponse Response Object
type ShowProgressDataResponse struct {

	// 总数。
	Count *int64 `json:"count,omitempty"`

	// 数据生成时间，UTC时间，例如：2020-09-01T18:50:20.200Z
	CreateTime *string `json:"create_time,omitempty"`

	// 对比结果。
	FlowCompareData *[]FlowCompareData `json:"flow_compare_data,omitempty"`
	HttpStatusCode  int                `json:"-"`
}

func (o ShowProgressDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProgressDataResponse struct{}"
	}

	return strings.Join([]string{"ShowProgressDataResponse", string(data)}, " ")
}
