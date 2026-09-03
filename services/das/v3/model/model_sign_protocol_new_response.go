package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SignProtocolNewResponse Response Object
type SignProtocolNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SignProtocolNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignProtocolNewResponse struct{}"
	}

	return strings.Join([]string{"SignProtocolNewResponse", string(data)}, " ")
}
