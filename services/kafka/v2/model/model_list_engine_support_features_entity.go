package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例支持的功能特性。
type ListEngineSupportFeaturesEntity struct {

	// 功能名称。
	Name *string `json:"name,omitempty" xml:"name"`

	Properties *ListEngineSupportFeaturesPropertiesEntity `json:"properties,omitempty" xml:"properties"`
}

func (o ListEngineSupportFeaturesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineSupportFeaturesEntity struct{}"
	}

	return strings.Join([]string{"ListEngineSupportFeaturesEntity", string(data)}, " ")
}
