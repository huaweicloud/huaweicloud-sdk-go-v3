package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobResourceRequest Request Object
type DeleteJobResourceRequest struct {

	// 资源名。
	ResourceName string `json:"resource_name"`

	// 资源所在分组
	Group *string `json:"group,omitempty"`
}

func (o DeleteJobResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobResourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteJobResourceRequest", string(data)}, " ")
}
