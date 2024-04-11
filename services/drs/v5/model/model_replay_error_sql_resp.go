package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplayErrorSqlResp 回放异常SQL数据项
type ReplayErrorSqlResp struct {

	// SQL类型
	ObjectType *string `json:"object_type,omitempty"`

	// SQL语句
	AbnormalSql *string `json:"abnormal_sql,omitempty"`

	// 异常原因描述
	AbnormalInfo *string `json:"abnormal_info,omitempty"`
}

func (o ReplayErrorSqlResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplayErrorSqlResp struct{}"
	}

	return strings.Join([]string{"ReplayErrorSqlResp", string(data)}, " ")
}
