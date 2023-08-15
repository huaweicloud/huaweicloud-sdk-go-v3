package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateNatResponse Response Object
type DeletePrivateNatResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateNatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateNatResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateNatResponse", string(data)}, " ")
}
