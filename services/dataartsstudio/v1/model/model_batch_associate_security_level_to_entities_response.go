package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAssociateSecurityLevelToEntitiesResponse Response Object
type BatchAssociateSecurityLevelToEntitiesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAssociateSecurityLevelToEntitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociateSecurityLevelToEntitiesResponse struct{}"
	}

	return strings.Join([]string{"BatchAssociateSecurityLevelToEntitiesResponse", string(data)}, " ")
}
