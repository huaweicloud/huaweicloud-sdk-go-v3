package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDirectConnectResponse Response Object
type DeleteDirectConnectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"DeleteDirectConnectResponse", string(data)}, " ")
}
