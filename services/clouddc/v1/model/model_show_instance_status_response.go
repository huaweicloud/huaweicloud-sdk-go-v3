package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceStatusResponse Response Object
type ShowInstanceStatusResponse struct {
	State *InstanceState `json:"state,omitempty"`

	Error          *ErrorStatus `json:"error,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowInstanceStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceStatusResponse", string(data)}, " ")
}
