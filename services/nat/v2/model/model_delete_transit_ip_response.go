package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTransitIpResponse Response Object
type DeleteTransitIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTransitIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransitIpResponse struct{}"
	}

	return strings.Join([]string{"DeleteTransitIpResponse", string(data)}, " ")
}
