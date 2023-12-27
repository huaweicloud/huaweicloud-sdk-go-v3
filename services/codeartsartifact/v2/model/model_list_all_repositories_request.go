package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllRepositoriesRequest Request Object
type ListAllRepositoriesRequest struct {

	// 租户id
	TenantId string `json:"tenant_id"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 组id
	GroupId *string `json:"group_id,omitempty"`

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 排序类型
	Sort *string `json:"sort,omitempty"`

	// 查询内容
	Qname *string `json:"qname,omitempty"`

	// 仓库类型
	Type *string `json:"type,omitempty"`

	// 仓库格式
	Format *string `json:"format,omitempty"`

	// 仓库格式列表
	FormatList *string `json:"format_list,omitempty"`

	// 是否是回收站文件
	IsRecycleBin *bool `json:"is_recycle_bin,omitempty"`

	// 是否需要分页
	IsNeedPaging *bool `json:"is_need_paging,omitempty"`

	// 是否在项目中
	InProject *bool `json:"in_project,omitempty"`
}

func (o ListAllRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAllRepositoriesRequest", string(data)}, " ")
}
