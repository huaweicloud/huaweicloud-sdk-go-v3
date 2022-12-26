package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchTemplateByIdResponse struct {
	ApproveInfo *ApproveInfo `json:"approve_info,omitempty"`

	// 模板创人，从接口调用传入的token中获取。
	CreateBy *string `json:"create_by,omitempty"`

	// 模板创建时间，为utc时间毫秒数。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 模板id，唯一标识，根据project_id和template_name生成。
	Id *string `json:"id,omitempty"`

	// 模板是否收藏，不允许更新模板时修改，更改收藏状态调用单独的更新模板收藏状态接口
	IsCollect *bool `json:"is_collect,omitempty"`

	// 是否发布成服务
	IsPublish *bool `json:"is_publish,omitempty"`

	// 作业id
	JobId *string `json:"job_id,omitempty"`

	// 作业版本
	JobVersion *int32 `json:"job_version,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 是否需要同步
	NeedSynchronize *bool `json:"need_synchronize,omitempty"`

	// 任务执行时需要的参数列表。
	Nodes *[]Node `json:"nodes,omitempty"`

	// 任务执行时需要的参数列表。
	Parameters *[]Parameter `json:"parameters,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 引用参数
	Quote *[]string `json:"quote,omitempty"`

	RateControl *RateControl `json:"rate_control,omitempty"`

	// 默认模板为public，自定义模板为private
	ShareType *string `json:"share_type,omitempty"`

	// 作业步骤
	Steps *[]Step `json:"steps,omitempty"`

	// 模板更新人，从接口调用传入的token中获取。
	UpdateBy *string `json:"update_by,omitempty"`

	// 模板更新时间，为utc时间毫秒数。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 模板版本
	Version        *string `json:"version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SearchTemplateByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchTemplateByIdResponse struct{}"
	}

	return strings.Join([]string{"SearchTemplateByIdResponse", string(data)}, " ")
}
