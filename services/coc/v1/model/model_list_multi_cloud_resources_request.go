package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMultiCloudResourcesRequest Request Object
type ListMultiCloudResourcesRequest struct {

	// 云厂商
	Vendor string `json:"vendor"`

	// 资源类型
	Type *string `json:"type,omitempty"`

	// 最大的返回数量
	Limit int32 `json:"limit"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	// 资源id列表
	ResourceIdList *[]string `json:"resource_id_list,omitempty"`

	// 资源名称
	NameList *[]string `json:"name_list,omitempty"`

	// region id列表
	RegionIdList *[]string `json:"region_id_list,omitempty"`
}

func (o ListMultiCloudResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMultiCloudResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListMultiCloudResourcesRequest", string(data)}, " ")
}
