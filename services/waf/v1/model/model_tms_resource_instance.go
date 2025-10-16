package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmsResourceInstance struct {
	ResourceId *string `json:"resource_id,omitempty"`

	ResourceName *string `json:"resource_name,omitempty"`

	ResourceDetail *interface{} `json:"resource_detail,omitempty"`

	Tags *[]TmsResourceTag `json:"tags,omitempty"`

	SysTags *[]TmsResourceTag `json:"sysTags,omitempty"`
}

func (o TmsResourceInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsResourceInstance struct{}"
	}

	return strings.Join([]string{"TmsResourceInstance", string(data)}, " ")
}
