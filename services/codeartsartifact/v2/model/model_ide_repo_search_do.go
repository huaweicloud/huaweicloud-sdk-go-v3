package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdeRepoSearchDo struct {

	// 搜索制品名称
	ArtifactName string `json:"artifact_name"`

	// 制品类型
	ArtifactType *string `json:"artifact_type,omitempty"`

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页条数
	PageSize *int32 `json:"page_size,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 是否在项目中
	InProject *string `json:"in_project,omitempty"`
}

func (o IdeRepoSearchDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdeRepoSearchDo struct{}"
	}

	return strings.Join([]string{"IdeRepoSearchDo", string(data)}, " ")
}
