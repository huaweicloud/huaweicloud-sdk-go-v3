package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlgorithmRequest Request Object
type CreateAlgorithmRequest struct {
	Body *Algorithm `json:"body,omitempty"`
}

func (o CreateAlgorithmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlgorithmRequest struct{}"
	}

	return strings.Join([]string{"CreateAlgorithmRequest", string(data)}, " ")
}
