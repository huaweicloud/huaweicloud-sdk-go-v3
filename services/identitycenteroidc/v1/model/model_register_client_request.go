package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterClientRequest Request Object
type RegisterClientRequest struct {
	Body *RegisterClientReqBody `json:"body,omitempty"`
}

func (o RegisterClientRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterClientRequest struct{}"
	}

	return strings.Join([]string{"RegisterClientRequest", string(data)}, " ")
}
