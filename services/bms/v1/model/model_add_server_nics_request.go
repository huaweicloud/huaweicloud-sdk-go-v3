package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddServerNicsRequest Request Object
type AddServerNicsRequest struct {
	ServerId string `json:"server_id"`

	Body *AddServerNicsReq `json:"body,omitempty"`
}

func (o AddServerNicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServerNicsRequest struct{}"
	}

	return strings.Join([]string{"AddServerNicsRequest", string(data)}, " ")
}
