package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInstancesResponse struct {

	// DDM实例的信息。
	Instances *[]ShowInstanceBeanResponse `json:"instances,omitempty" xml:"instances"`

	// 租户下的DDM实例个数
	InstanceNum *int32 `json:"instance_num,omitempty" xml:"instance_num"`

	// 当前页码
	PageNo *int32 `json:"page_no,omitempty" xml:"page_no"`

	// 当前页码的数据条数
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size"`

	// 总条数
	TotalRecord *int32 `json:"total_record,omitempty" xml:"total_record"`

	// 总页数
	TotalPage      *int32 `json:"total_page,omitempty" xml:"total_page"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesResponse", string(data)}, " ")
}
