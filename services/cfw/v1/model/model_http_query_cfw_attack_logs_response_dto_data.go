package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpQueryCfwAttackLogsResponseDtoData 查询攻击日志返回值
type HttpQueryCfwAttackLogsResponseDtoData struct {

	// 返回数量
	Total *int32 `json:"total,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 记录
	Records *[]HttpQueryCfwAttackLogsResponseDtoDataRecords `json:"records,omitempty"`
}

func (o HttpQueryCfwAttackLogsResponseDtoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpQueryCfwAttackLogsResponseDtoData struct{}"
	}

	return strings.Join([]string{"HttpQueryCfwAttackLogsResponseDtoData", string(data)}, " ")
}
