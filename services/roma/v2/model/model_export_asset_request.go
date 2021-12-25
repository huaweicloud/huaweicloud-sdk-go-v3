package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportAssetRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *AssetExportRequest `json:"body,omitempty"`
}

func (o ExportAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAssetRequest struct{}"
	}

	return strings.Join([]string{"ExportAssetRequest", string(data)}, " ")
}
