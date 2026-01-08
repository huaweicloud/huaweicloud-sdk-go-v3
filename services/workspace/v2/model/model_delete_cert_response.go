package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCertResponse Response Object
type DeleteCertResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertResponse struct{}"
	}

	return strings.Join([]string{"DeleteCertResponse", string(data)}, " ")
}
