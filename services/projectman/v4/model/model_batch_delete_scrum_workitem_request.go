package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteScrumWorkitemRequest Request Object
type BatchDeleteScrumWorkitemRequest struct {
	Body *BatchDeleteModuleRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteScrumWorkitemRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScrumWorkitemRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteScrumWorkitemRequest", string(data)}, " ")
}
