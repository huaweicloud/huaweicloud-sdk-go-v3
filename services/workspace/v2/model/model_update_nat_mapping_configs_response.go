package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNatMappingConfigsResponse Response Object
type UpdateNatMappingConfigsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNatMappingConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatMappingConfigsResponse struct{}"
	}

	return strings.Join([]string{"UpdateNatMappingConfigsResponse", string(data)}, " ")
}
