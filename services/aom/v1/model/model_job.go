package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 作业id
type Job struct {

	// 作业id。
	Id *string `json:"id,omitempty"`

	// 作业名称。
	Name string `json:"name"`

	// 实体的创建时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 实体的最后更新时间戳。 注意：执行创建/修补/删除操作时，update_time将更新。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 修改人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 作业描述，最大长度为1000。
	Description *string `json:"description,omitempty"`

	// 企业项目id。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 租户从IAM申请到的projectid，一般为32位字符串。
	ProjectId string `json:"project_id"`

	// 作业步骤。
	Steps []Step `json:"steps"`

	// 全局参数。
	Parameters []Parameter `json:"parameters"`

	RateControl *RateControl `json:"rate_control"`

	ApproveInfo *ApproveInfo `json:"approve_info"`

	// 是否为最新版本的作业
	IsLatestVersion *bool `json:"is_latest_version,omitempty"`

	// 版本号
	VersionNumber *int32 `json:"version_number,omitempty"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}
