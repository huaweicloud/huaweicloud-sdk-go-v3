package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpgradeDatabaseVersionRequest Request Object
type BatchUpgradeDatabaseVersionRequest struct {
	Body *BatchUpgradeDatabaseVersionRequestBody `json:"body,omitempty"`
}

func (o BatchUpgradeDatabaseVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpgradeDatabaseVersionRequest struct{}"
	}

	return strings.Join([]string{"BatchUpgradeDatabaseVersionRequest", string(data)}, " ")
}
