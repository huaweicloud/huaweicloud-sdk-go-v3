package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmsResourceInstance struct {

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源细节
	ResourceDetail *interface{} `json:"resource_detail,omitempty"`

	// 系统标签键值对
	Tags *[]TmsResourceTag `json:"tags,omitempty"`

	// 系统标签键值对
	SysTags *[]TmsResourceTag `json:"sysTags,omitempty"`
}

func (o TmsResourceInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsResourceInstance struct{}"
	}

	return strings.Join([]string{"TmsResourceInstance", string(data)}, " ")
}
