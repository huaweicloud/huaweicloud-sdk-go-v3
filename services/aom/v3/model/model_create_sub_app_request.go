package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubAppRequest Request Object
type CreateSubAppRequest struct {
	Body *SubAppCreateParam `json:"body,omitempty"`
}

func (o CreateSubAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubAppRequest struct{}"
	}

	return strings.Join([]string{"CreateSubAppRequest", string(data)}, " ")
}
