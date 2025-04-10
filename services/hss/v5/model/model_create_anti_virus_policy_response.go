package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAntiVirusPolicyResponse Response Object
type CreateAntiVirusPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAntiVirusPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAntiVirusPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateAntiVirusPolicyResponse", string(data)}, " ")
}
