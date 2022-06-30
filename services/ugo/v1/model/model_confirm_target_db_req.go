package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 确认目标数据库版本的请求体。
type ConfirmTargetDbReq struct {

	// 项目ID。
	EvaluationProjectId string `json:"evaluation_project_id"`

	// 目标库类型。
	TargetDbType string `json:"target_db_type"`

	// 目标库版本。
	TargetDbVersion string `json:"target_db_version"`
}

func (o ConfirmTargetDbReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmTargetDbReq struct{}"
	}

	return strings.Join([]string{"ConfirmTargetDbReq", string(data)}, " ")
}
