package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDedicatedResourceInfoRequest Request Object
type ShowDedicatedResourceInfoRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 专属资源池ID。
	DedicatedResourceId string `json:"dedicated_resource_id"`
}

func (o ShowDedicatedResourceInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDedicatedResourceInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowDedicatedResourceInfoRequest", string(data)}, " ")
}
