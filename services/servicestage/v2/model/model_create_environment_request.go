package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEnvironmentRequest struct {
	Body *EnvironmentCreate `json:"body,omitempty" xml:"body"`
}

func (o CreateEnvironmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentRequest struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentRequest", string(data)}, " ")
}
