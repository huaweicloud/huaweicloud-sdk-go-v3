package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrsTaskReq 创建drs任务的请求体
type CreateDrsTaskReq struct {

	// 目标实例id
	TargetInstanceId string `json:"target_instance_id"`

	// 目标实例用户账号
	TargetUserName string `json:"target_user_name"`

	// 目标实例用户密码
	TargetUserPassword string `json:"target_user_password"`

	// 源实例用户账号
	SourceUserName string `json:"source_user_name"`

	// 源实例用户密码
	SourceUserPassword string `json:"source_user_password"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// Drs实例规格
	DrsNodeType string `json:"drs_node_type"`

	// 数据库名称
	DatabaseList []string `json:"database_list"`
}

func (o CreateDrsTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrsTaskReq struct{}"
	}

	return strings.Join([]string{"CreateDrsTaskReq", string(data)}, " ")
}
