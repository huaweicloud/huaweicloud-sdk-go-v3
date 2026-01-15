package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateDomainControllerResponse Response Object
type ValidateDomainControllerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ValidateDomainControllerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateDomainControllerResponse struct{}"
	}

	return strings.Join([]string{"ValidateDomainControllerResponse", string(data)}, " ")
}
