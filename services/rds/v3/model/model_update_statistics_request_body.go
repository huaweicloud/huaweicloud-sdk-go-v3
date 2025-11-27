package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStatisticsRequestBody 统计信息更新内容
type UpdateStatisticsRequestBody struct {

	// 数据库名
	DbName *string `json:"db_name,omitempty"`
}

func (o UpdateStatisticsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStatisticsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateStatisticsRequestBody", string(data)}, " ")
}
