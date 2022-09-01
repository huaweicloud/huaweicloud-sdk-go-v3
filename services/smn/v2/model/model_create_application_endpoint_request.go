package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateApplicationEndpointRequest struct {

	// Application的唯一资源标识，可通过[查询Application](https://support.huaweicloud.com/api-smn/ListApplications.html)获取该标识。
	ApplicationUrn string `json:"application_urn" xml:"application_urn"`

	Body *CreateApplicationEndpointRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateApplicationEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationEndpointRequest struct{}"
	}

	return strings.Join([]string{"CreateApplicationEndpointRequest", string(data)}, " ")
}
