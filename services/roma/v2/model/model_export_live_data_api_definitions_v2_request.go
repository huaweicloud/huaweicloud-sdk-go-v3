package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportLiveDataApiDefinitionsV2Request Request Object
type ExportLiveDataApiDefinitionsV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *LiveDataExportReq `json:"body,omitempty"`
}

func (o ExportLiveDataApiDefinitionsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportLiveDataApiDefinitionsV2Request struct{}"
	}

	return strings.Join([]string{"ExportLiveDataApiDefinitionsV2Request", string(data)}, " ")
}
