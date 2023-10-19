package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociatePrincipalsResponse Response Object
type AssociatePrincipalsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociatePrincipalsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatePrincipalsResponse struct{}"
	}

	return strings.Join([]string{"AssociatePrincipalsResponse", string(data)}, " ")
}
