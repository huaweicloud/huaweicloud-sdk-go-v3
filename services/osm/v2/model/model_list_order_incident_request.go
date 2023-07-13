package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrderIncidentRequest Request Object
type ListOrderIncidentRequest struct {

	// 版本号
	Version *string `json:"version,omitempty"`

	// 关键字
	SearchKey *[]string `json:"searchKey,omitempty"`

	// 标签列表
	LabelIdList *[]int32 `json:"labelIdList,omitempty"`

	// appKey
	AppKey *string `json:"appKey,omitempty"`

	// incidentId
	IncidentId *string `json:"incidentId,omitempty"`

	// 查询开始时间
	QueryStartTime *string `json:"queryStartTime,omitempty"`

	// 查询结束时间
	QueryEndTime *string `json:"queryEndTime,omitempty"`

	// 状态
	Status *int32 `json:"status,omitempty"`

	// 工单状态
	IncidentStatus *string `json:"incidentStatus,omitempty"`

	// 用户名称
	XCustomerName *string `json:"xCustomerName,omitempty"`

	// 分组
	GroupId *string `json:"groupId,omitempty"`

	// 产品分类
	ProductCategoryId *string `json:"productCategoryId,omitempty"`

	// 类型
	BusinessTypeId *string `json:"businessTypeId,omitempty"`

	// 页码
	PageNo *int32 `json:"pageNo,omitempty"`

	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`

	// 客户id
	XCustomerId *string `json:"xCustomerId,omitempty"`
}

func (o ListOrderIncidentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrderIncidentRequest struct{}"
	}

	return strings.Join([]string{"ListOrderIncidentRequest", string(data)}, " ")
}
