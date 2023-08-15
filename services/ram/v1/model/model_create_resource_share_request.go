package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceShareRequest Request Object
type CreateResourceShareRequest struct {
	Body *CreateResourceShareReqBody `json:"body,omitempty"`
}

func (o CreateResourceShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceShareRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceShareRequest", string(data)}, " ")
}
