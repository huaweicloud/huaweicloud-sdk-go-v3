package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResultsTopnBody struct {

	// 索引流量，byte, 查询数据类型中包含index时返回
	IndexTraffic float64 `json:"index_traffic"`

	// 存储量，byte, 查询数据类型中包含storage时返回
	Storage float64 `json:"storage"`

	// 写入流量，byte, 查询数据类型中包含write时返回
	WriteTraffic float64 `json:"write_traffic"`

	// 日志组id，资源类型为日志组时返回
	LogGroupId string `json:"log_group_id"`

	// 日志组名称，资源类型为日志组时返回
	LogGroupName string `json:"log_group_name"`

	// 日志流id，资源类型为日志流时返回
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志流名称，资源类型为日志流时返回
	LogStreamName *string `json:"log_stream_name,omitempty"`
}

func (o ResultsTopnBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultsTopnBody struct{}"
	}

	return strings.Join([]string{"ResultsTopnBody", string(data)}, " ")
}
