package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertResponse Response Object
type CreateCertResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertResponse struct{}"
	}

	return strings.Join([]string{"CreateCertResponse", string(data)}, " ")
}
