package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadSchemaMetadataRequest Request Object
type UploadSchemaMetadataRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	Body *LoadSchemaMetadataReq `json:"body,omitempty"`
}

func (o UploadSchemaMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSchemaMetadataRequest struct{}"
	}

	return strings.Join([]string{"UploadSchemaMetadataRequest", string(data)}, " ")
}
