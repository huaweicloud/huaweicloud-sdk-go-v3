package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiEnablement struct {

	// 资源组版本
	GroupVersion *string `json:"groupVersion,omitempty"`

	// 资源类型、名称
	Resources *[]ApiResource `json:"resources,omitempty"`
}

func (o ApiEnablement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiEnablement struct{}"
	}

	return strings.Join([]string{"ApiEnablement", string(data)}, " ")
}
