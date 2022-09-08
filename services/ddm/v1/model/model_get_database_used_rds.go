package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// used_rds 返回参数
type GetDatabaseUsedRds struct {

	// 关联RDS节点ID。
	Id string `json:"id"`

	// 关联RDS名称
	Name string `json:"name"`

	// 关联RDS状态。
	Status string `json:"status"`

	// 响应信息，若无异常信息则不返回该参数。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o GetDatabaseUsedRds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDatabaseUsedRds struct{}"
	}

	return strings.Join([]string{"GetDatabaseUsedRds", string(data)}, " ")
}
