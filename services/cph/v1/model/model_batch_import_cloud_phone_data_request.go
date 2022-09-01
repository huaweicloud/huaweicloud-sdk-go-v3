package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchImportCloudPhoneDataRequest struct {
	Body *BatchImportCloudPhoneDataRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchImportCloudPhoneDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportCloudPhoneDataRequest struct{}"
	}

	return strings.Join([]string{"BatchImportCloudPhoneDataRequest", string(data)}, " ")
}
