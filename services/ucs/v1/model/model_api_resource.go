package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiResource struct {

	// 资源名
	Name *string `json:"name,omitempty"`

	// 资源类别
	Kind *string `json:"kind,omitempty"`
}

func (o ApiResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiResource struct{}"
	}

	return strings.Join([]string{"ApiResource", string(data)}, " ")
}
