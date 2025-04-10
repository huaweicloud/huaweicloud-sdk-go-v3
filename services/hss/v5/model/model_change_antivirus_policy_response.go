package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAntivirusPolicyResponse Response Object
type ChangeAntivirusPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeAntivirusPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAntivirusPolicyResponse struct{}"
	}

	return strings.Join([]string{"ChangeAntivirusPolicyResponse", string(data)}, " ")
}
