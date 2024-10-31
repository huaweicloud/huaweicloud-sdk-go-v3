package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpQueryCfwAccessControllerLogsResponseDtoData 查询访问控制日志返回数据
type HttpQueryCfwAccessControllerLogsResponseDtoData struct {

	// 查询访问控制日志记录总数
	Total *int32 `json:"total,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 查询访问控制日志记录
	Records *[]HttpQueryCfwAccessControllerLogsResponseDtoDataRecords `json:"records,omitempty"`
}

func (o HttpQueryCfwAccessControllerLogsResponseDtoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpQueryCfwAccessControllerLogsResponseDtoData struct{}"
	}

	return strings.Join([]string{"HttpQueryCfwAccessControllerLogsResponseDtoData", string(data)}, " ")
}
