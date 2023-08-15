package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateSnatResponse Response Object
type DeletePrivateSnatResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateSnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateSnatResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateSnatResponse", string(data)}, " ")
}
