package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadSchemaMetadataRequest Request Object
type DownloadSchemaMetadataRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`
}

func (o DownloadSchemaMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSchemaMetadataRequest struct{}"
	}

	return strings.Join([]string{"DownloadSchemaMetadataRequest", string(data)}, " ")
}
