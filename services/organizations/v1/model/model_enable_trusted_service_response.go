package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type EnableTrustedServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableTrustedServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableTrustedServiceResponse struct{}"
	}

	return strings.Join([]string{"EnableTrustedServiceResponse", string(data)}, " ")
}
