package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterAuthorizedQueueRequest Request Object
type RegisterAuthorizedQueueRequest struct {
	Body *GrantQueuePermissionReq `json:"body,omitempty"`
}

func (o RegisterAuthorizedQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterAuthorizedQueueRequest struct{}"
	}

	return strings.Join([]string{"RegisterAuthorizedQueueRequest", string(data)}, " ")
}
