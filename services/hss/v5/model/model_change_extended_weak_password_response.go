package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeExtendedWeakPasswordResponse Response Object
type ChangeExtendedWeakPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeExtendedWeakPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeExtendedWeakPasswordResponse struct{}"
	}

	return strings.Join([]string{"ChangeExtendedWeakPasswordResponse", string(data)}, " ")
}
