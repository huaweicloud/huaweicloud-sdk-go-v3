package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateStructConfigRequest struct {
	Body *StructConfig `json:"body,omitempty"`
}

func (o UpdateStructConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStructConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateStructConfigRequest", string(data)}, " ")
}
