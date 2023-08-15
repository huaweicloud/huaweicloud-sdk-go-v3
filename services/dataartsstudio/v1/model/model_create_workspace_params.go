package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkspaceParams IAM用户组信息
type CreateWorkspaceParams struct {

	// DLI脏数据OBS路径
	BadRecordLocationName *string `json:"bad_record_location_name,omitempty"`

	// 工作空间描述
	Description *string `json:"description,omitempty"`

	// 企业项目id，如果当前为公有云，且用户开启企业项目，则必选
	EpsId string `json:"eps_id"`

	// 作业日志OBS路径
	JobLogLocationName *string `json:"job_log_location_name,omitempty"`

	// 工作空间名称
	Name string `json:"name"`
}

func (o CreateWorkspaceParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkspaceParams struct{}"
	}

	return strings.Join([]string{"CreateWorkspaceParams", string(data)}, " ")
}
