package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchMigrateCloudPhoneRequest struct {
	Body *BatchMigrateCloudPhoneRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchMigrateCloudPhoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMigrateCloudPhoneRequest struct{}"
	}

	return strings.Join([]string{"BatchMigrateCloudPhoneRequest", string(data)}, " ")
}
