package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerNicsRequest Request Object
type DeleteServerNicsRequest struct {
	ServerId string `json:"server_id"`

	Body *DeleteServerNicsReq `json:"body,omitempty"`
}

func (o DeleteServerNicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerNicsRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerNicsRequest", string(data)}, " ")
}
