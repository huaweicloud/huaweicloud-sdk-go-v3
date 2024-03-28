package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkJobResource 创建会话请求参数resources的元素。
type SparkJobResource struct {

	// 资源名称。
	Name *string `json:"name,omitempty"`

	// 资源类型。
	Type *string `json:"type,omitempty"`
}

func (o SparkJobResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJobResource struct{}"
	}

	return strings.Join([]string{"SparkJobResource", string(data)}, " ")
}
