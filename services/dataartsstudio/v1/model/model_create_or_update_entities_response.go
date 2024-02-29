package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateEntitiesResponse Response Object
type CreateOrUpdateEntitiesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateOrUpdateEntitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateEntitiesResponse struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateEntitiesResponse", string(data)}, " ")
}
