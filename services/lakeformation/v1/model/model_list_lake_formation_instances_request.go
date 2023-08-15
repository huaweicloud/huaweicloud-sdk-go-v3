package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLakeFormationInstancesRequest Request Object
type ListLakeFormationInstancesRequest struct {

	// 是否查询回收站中的实例
	InRecycleBin bool `json:"in_recycle_bin"`

	// 分页查询时的偏移量
	Offset int32 `json:"offset"`

	// 分页一页显示数
	Limit int32 `json:"limit"`

	// 使用LakeFormation实例名进行检索
	Name *string `json:"name,omitempty"`

	// 企业项目id
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 标签条件列表
	Tags *string `json:"tags,omitempty"`
}

func (o ListLakeFormationInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLakeFormationInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListLakeFormationInstancesRequest", string(data)}, " ")
}
