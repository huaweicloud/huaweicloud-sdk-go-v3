package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DependPackage 依赖包信息
type DependPackage struct {

	// 文件类型
	Type *string `json:"type,omitempty"`

	// 文件路径
	Location *string `json:"location,omitempty"`
}

func (o DependPackage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DependPackage struct{}"
	}

	return strings.Join([]string{"DependPackage", string(data)}, " ")
}
