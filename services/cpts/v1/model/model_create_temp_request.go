package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTempRequest struct {
	Body *CreateTempRequestBody `json:"body,omitempty"`
}

func (o CreateTempRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTempRequest struct{}"
	}

	return strings.Join([]string{"CreateTempRequest", string(data)}, " ")
}
