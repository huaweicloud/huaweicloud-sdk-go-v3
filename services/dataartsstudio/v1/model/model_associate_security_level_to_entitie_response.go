package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateSecurityLevelToEntitieResponse Response Object
type AssociateSecurityLevelToEntitieResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateSecurityLevelToEntitieResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSecurityLevelToEntitieResponse struct{}"
	}

	return strings.Join([]string{"AssociateSecurityLevelToEntitieResponse", string(data)}, " ")
}
