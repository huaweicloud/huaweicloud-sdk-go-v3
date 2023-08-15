package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLabelRequest Request Object
type CreateLabelRequest struct {
	Body *CreateLabelReq `json:"body,omitempty"`
}

func (o CreateLabelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLabelRequest struct{}"
	}

	return strings.Join([]string{"CreateLabelRequest", string(data)}, " ")
}
