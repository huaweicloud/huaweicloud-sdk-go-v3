package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBusinessRequest Request Object
type CreateBusinessRequest struct {
	Body *CreateBusinessModel `json:"body,omitempty"`
}

func (o CreateBusinessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBusinessRequest struct{}"
	}

	return strings.Join([]string{"CreateBusinessRequest", string(data)}, " ")
}
