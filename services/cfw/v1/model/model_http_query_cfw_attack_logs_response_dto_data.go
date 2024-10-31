package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpQueryCfwAttackLogsResponseDtoData 查询攻击日志返回值数据
type HttpQueryCfwAttackLogsResponseDtoData struct {

	// 返回攻击数据总数
	Total *int32 `json:"total,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 攻击日志记录列表
	Records *[]HttpQueryCfwAttackLogsResponseDtoDataRecords `json:"records,omitempty"`
}

func (o HttpQueryCfwAttackLogsResponseDtoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpQueryCfwAttackLogsResponseDtoData struct{}"
	}

	return strings.Join([]string{"HttpQueryCfwAttackLogsResponseDtoData", string(data)}, " ")
}
