package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询流日志返回值
type HttpQueryCfwFlowLogsResponseDtoData struct {

	// 返回数量
	Total *int32 `json:"total,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 记录
	Records *[]HttpQueryCfwFlowLogsResponseDtoDataRecords `json:"records,omitempty"`
}

func (o HttpQueryCfwFlowLogsResponseDtoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpQueryCfwFlowLogsResponseDtoData struct{}"
	}

	return strings.Join([]string{"HttpQueryCfwFlowLogsResponseDtoData", string(data)}, " ")
}
