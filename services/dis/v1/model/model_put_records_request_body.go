package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PutRecordsRequestBody struct {
	// 已创建的通道名称。

	StreamName string `json:"stream_name"`
	// 通道唯一标识符。  当使用stream_name没有找到对应通道且stream_id不为空时，会使用stream_id去查找通道。  说明：  上传数据到被授权的通道时，必须配置此参数。

	StreamId *string `json:"stream_id,omitempty"`

	Records []PutRecordsRequestEntry `json:"records"`
}

func (o PutRecordsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutRecordsRequestBody struct{}"
	}

	return strings.Join([]string{"PutRecordsRequestBody", string(data)}, " ")
}
