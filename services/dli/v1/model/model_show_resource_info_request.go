package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceInfoRequest Request Object
type ShowResourceInfoRequest struct {

	// 资源名。
	ResourceName string `json:"resource_name"`

	Group *string `json:"group,omitempty"`
}

func (o ShowResourceInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceInfoRequest", string(data)}, " ")
}
