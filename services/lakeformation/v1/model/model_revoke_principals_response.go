package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevokePrincipalsResponse Response Object
type RevokePrincipalsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RevokePrincipalsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokePrincipalsResponse struct{}"
	}

	return strings.Join([]string{"RevokePrincipalsResponse", string(data)}, " ")
}
