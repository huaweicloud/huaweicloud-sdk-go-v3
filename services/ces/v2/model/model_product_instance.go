package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductInstance struct {

	// 资源首层维度，如：弹性云服务器，则维度为instance_id；”。
	FirstDimensionName string `json:"first_dimension_name"`

	// 资源首层维度值，为资源的实例ID，如：4270ff17-aba3-4138-89fa-820594c39755。
	FirstDimensionValue string `json:"first_dimension_value"`

	// 资源名称
	ResourceName string `json:"resource_name"`
}

func (o ProductInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInstance struct{}"
	}

	return strings.Join([]string{"ProductInstance", string(data)}, " ")
}
