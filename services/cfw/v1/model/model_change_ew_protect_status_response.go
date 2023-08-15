package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeEwProtectStatusResponse Response Object
type ChangeEwProtectStatusResponse struct {
	Data *SuccessRspData `json:"data,omitempty"`

	// trace id
	TraceId        *string `json:"trace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeEwProtectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEwProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeEwProtectStatusResponse", string(data)}, " ")
}
