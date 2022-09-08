package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchExportCloudPhoneDataRequest struct {
	Body *BatchExportCloudPhoneDataRequestBody `json:"body,omitempty"`
}

func (o BatchExportCloudPhoneDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExportCloudPhoneDataRequest struct{}"
	}

	return strings.Join([]string{"BatchExportCloudPhoneDataRequest", string(data)}, " ")
}
