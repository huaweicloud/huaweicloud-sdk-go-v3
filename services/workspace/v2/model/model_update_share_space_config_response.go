package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateShareSpaceConfigResponse Response Object
type UpdateShareSpaceConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateShareSpaceConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateShareSpaceConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateShareSpaceConfigResponse", string(data)}, " ")
}
