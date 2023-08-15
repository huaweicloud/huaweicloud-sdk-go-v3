package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppServersRequest Request Object
type CreateAppServersRequest struct {
	Body *CreateAppServerReq `json:"body,omitempty"`
}

func (o CreateAppServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppServersRequest struct{}"
	}

	return strings.Join([]string{"CreateAppServersRequest", string(data)}, " ")
}
