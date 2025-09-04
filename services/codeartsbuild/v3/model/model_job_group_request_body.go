package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobGroupRequestBody 创建更新分组返回体
type JobGroupRequestBody struct {

	// 主键id
	Id *string `json:"id,omitempty"`

	// 分组名称
	Name *string `json:"name,omitempty"`

	// CodeArts项目ID。获取方式请参考[获取CodeArts项目ID](https://support.huaweicloud.com/api-codeci/cloudbuild_03_0022.html)。
	ProjectId *string `json:"project_id,omitempty"`

	// 父分组id
	ParentId *string `json:"parent_id,omitempty"`

	// 任务分组id
	GroupId *string `json:"group_id,omitempty"`
}

func (o JobGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobGroupRequestBody struct{}"
	}

	return strings.Join([]string{"JobGroupRequestBody", string(data)}, " ")
}
