package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkspaceDto 工作空间信息
type WorkspaceDto struct {

	// DLI脏数据OBS路径
	BadRecordLocationName *string `json:"bad_record_location_name,omitempty"`

	// 工作空间描述
	Description *string `json:"description,omitempty"`

	// 作业日志OBS路径
	JobLogLocationName *string `json:"job_log_location_name,omitempty"`

	// 工作空间名称
	Name string `json:"name"`

	// 企业项目ID
	EpsId *string `json:"eps_id,omitempty"`

	// 工作空间模式。0：简易模式；1：企业模式
	Mode *string `json:"mode,omitempty"`

	// 工作空间状态。0：正常；5：冻结
	Status *string `json:"status,omitempty"`
}

func (o WorkspaceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkspaceDto struct{}"
	}

	return strings.Join([]string{"WorkspaceDto", string(data)}, " ")
}
