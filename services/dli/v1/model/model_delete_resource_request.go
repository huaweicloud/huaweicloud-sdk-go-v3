package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceRequest Request Object
type DeleteResourceRequest struct {

	// 资源名。
	ResourceName string `json:"resource_name"`

	// 资源所在分组
	Group *string `json:"group,omitempty"`
}

func (o DeleteResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceRequest", string(data)}, " ")
}
