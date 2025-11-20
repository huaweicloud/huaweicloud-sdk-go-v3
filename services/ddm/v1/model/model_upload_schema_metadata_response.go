package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadSchemaMetadataResponse Response Object
type UploadSchemaMetadataResponse struct {

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 工作id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadSchemaMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSchemaMetadataResponse struct{}"
	}

	return strings.Join([]string{"UploadSchemaMetadataResponse", string(data)}, " ")
}
