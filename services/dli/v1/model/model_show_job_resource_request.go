package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobResourceRequest Request Object
type ShowJobResourceRequest struct {

	// 资源名。
	ResourceName string `json:"resource_name"`

	Group *string `json:"group,omitempty"`
}

func (o ShowJobResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResourceRequest struct{}"
	}

	return strings.Join([]string{"ShowJobResourceRequest", string(data)}, " ")
}
