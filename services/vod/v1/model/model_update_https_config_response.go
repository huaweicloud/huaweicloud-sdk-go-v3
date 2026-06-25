package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpsConfigResponse Response Object
type UpdateHttpsConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateHttpsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpsConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateHttpsConfigResponse", string(data)}, " ")
}
