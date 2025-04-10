package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAntivirusPolicyResponse Response Object
type DeleteAntivirusPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAntivirusPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAntivirusPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAntivirusPolicyResponse", string(data)}, " ")
}
