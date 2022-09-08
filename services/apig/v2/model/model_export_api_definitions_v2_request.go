package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportApiDefinitionsV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ExportOpenApiReq `json:"body,omitempty"`
}

func (o ExportApiDefinitionsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportApiDefinitionsV2Request struct{}"
	}

	return strings.Join([]string{"ExportApiDefinitionsV2Request", string(data)}, " ")
}
