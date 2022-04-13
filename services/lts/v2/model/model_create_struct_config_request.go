package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateStructConfigRequest struct {
	Body *StructConfig `json:"body,omitempty"`
}

func (o CreateStructConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStructConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateStructConfigRequest", string(data)}, " ")
}
