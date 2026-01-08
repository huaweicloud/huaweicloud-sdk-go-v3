package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetCertStatusResponse Response Object
type SetCertStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetCertStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetCertStatusResponse struct{}"
	}

	return strings.Join([]string{"SetCertStatusResponse", string(data)}, " ")
}
