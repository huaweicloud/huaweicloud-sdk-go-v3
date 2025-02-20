package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchListMultiCloudResourceResponseData struct {

	// CMDB生成uuid
	Id *string `json:"id,omitempty"`

	// 在云厂商上存的资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 云厂商
	Vendor *string `json:"vendor,omitempty"`

	// 资源类型
	Type *string `json:"type,omitempty"`

	// 云厂商账户id
	Datasource *string `json:"datasource,omitempty"`

	// regionId
	RegionId *string `json:"region_id,omitempty"`

	// 资源详细属性
	Properties map[string]interface{} `json:"properties,omitempty"`

	// 资源创建时间
	CreateTime *string `json:"create_time,omitempty"`
}

func (o BatchListMultiCloudResourceResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListMultiCloudResourceResponseData struct{}"
	}

	return strings.Join([]string{"BatchListMultiCloudResourceResponseData", string(data)}, " ")
}
