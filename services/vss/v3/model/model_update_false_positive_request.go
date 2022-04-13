package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateFalsePositiveRequest struct {
	Body *UpdateFalsePositiveRequestBody `json:"body,omitempty"`
}

func (o UpdateFalsePositiveRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFalsePositiveRequest struct{}"
	}

	return strings.Join([]string{"UpdateFalsePositiveRequest", string(data)}, " ")
}
