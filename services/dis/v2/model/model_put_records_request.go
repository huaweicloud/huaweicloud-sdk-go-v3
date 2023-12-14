package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PutRecordsRequest struct {

	// 已创建的通道名称。
	StreamName string `json:"stream_name"`

	// 通道唯一标识符。  当使用stream_name没有找到对应通道且stream_id不为空时，会使用stream_id去查找通道。  说明：  上传数据到被授权的通道时，必须配置此参数。
	StreamId *string `json:"stream_id,omitempty"`

	// 待上传的记录列表。
	Records []PutRecordsRequestEntry `json:"records"`
}

func (o PutRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutRecordsRequest struct{}"
	}

	return strings.Join([]string{"PutRecordsRequest", string(data)}, " ")
}
