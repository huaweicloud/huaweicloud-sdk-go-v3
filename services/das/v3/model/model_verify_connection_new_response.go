package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyConnectionNewResponse Response Object
type VerifyConnectionNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o VerifyConnectionNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyConnectionNewResponse struct{}"
	}

	return strings.Join([]string{"VerifyConnectionNewResponse", string(data)}, " ")
}
