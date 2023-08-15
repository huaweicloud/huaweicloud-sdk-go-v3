package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePortInfoResponse Response Object
type DeletePortInfoResponse struct {
	Data           *DeletePortResponseModel `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o DeletePortInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePortInfoResponse struct{}"
	}

	return strings.Join([]string{"DeletePortInfoResponse", string(data)}, " ")
}
