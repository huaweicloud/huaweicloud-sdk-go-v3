package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateDnatResponse Response Object
type DeletePrivateDnatResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateDnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateDnatResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateDnatResponse", string(data)}, " ")
}
