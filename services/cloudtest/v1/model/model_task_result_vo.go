package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskResultVo struct {

	// 结果URI
	Uri *string `json:"uri,omitempty"`

	// 测试套名称
	Name *string `json:"name,omitempty"`

	// 任务uri
	TaskUri *string `json:"task_uri,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 责任人
	Owner *string `json:"owner,omitempty"`

	// 责任人名称
	OwnerName *string `json:"owner_name,omitempty"`

	// 发布版本号
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 分支/迭代uri
	VersionUri *string `json:"version_uri,omitempty"`

	// 创建人id
	Creator *string `json:"creator,omitempty"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新人
	Updator *string `json:"updator,omitempty"`

	// 更新人名称
	UpdatorName *string `json:"updator_name,omitempty"`

	// 更新人名称
	UpdateTime *string `json:"update_time,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 测试结果名称
	ResultName *string `json:"result_name,omitempty"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`
}

func (o TaskResultVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskResultVo struct{}"
	}

	return strings.Join([]string{"TaskResultVo", string(data)}, " ")
}
