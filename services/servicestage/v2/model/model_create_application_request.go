package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationRequest Request Object
type CreateApplicationRequest struct {
	Body *ApplicationCreate `json:"body,omitempty"`
}

func (o CreateApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationRequest struct{}"
	}

	return strings.Join([]string{"CreateApplicationRequest", string(data)}, " ")
}
