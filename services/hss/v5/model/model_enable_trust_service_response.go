package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableTrustServiceResponse Response Object
type EnableTrustServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableTrustServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableTrustServiceResponse struct{}"
	}

	return strings.Join([]string{"EnableTrustServiceResponse", string(data)}, " ")
}
