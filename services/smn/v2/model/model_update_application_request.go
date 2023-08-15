package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationRequest Request Object
type UpdateApplicationRequest struct {

	// Application的唯一资源标识，可通过[查询Application](smn_api_57004.xml)获取该标识。
	ApplicationUrn string `json:"application_urn"`

	Body *UpdateApplicationRequestBody `json:"body,omitempty"`
}

func (o UpdateApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationRequest struct{}"
	}

	return strings.Join([]string{"UpdateApplicationRequest", string(data)}, " ")
}
