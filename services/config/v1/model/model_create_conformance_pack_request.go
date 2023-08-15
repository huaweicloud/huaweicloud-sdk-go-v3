package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConformancePackRequest Request Object
type CreateConformancePackRequest struct {
	Body *ConformancePackRequestBody `json:"body,omitempty"`
}

func (o CreateConformancePackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConformancePackRequest struct{}"
	}

	return strings.Join([]string{"CreateConformancePackRequest", string(data)}, " ")
}
