package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchMigrateCloudPhoneRequest Request Object
type BatchMigrateCloudPhoneRequest struct {
	Body *BatchMigrateCloudPhoneRequestBody `json:"body,omitempty"`
}

func (o BatchMigrateCloudPhoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMigrateCloudPhoneRequest struct{}"
	}

	return strings.Join([]string{"BatchMigrateCloudPhoneRequest", string(data)}, " ")
}
