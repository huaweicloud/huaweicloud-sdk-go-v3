package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateConfigResponse Response Object
type ValidateConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ValidateConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateConfigResponse struct{}"
	}

	return strings.Join([]string{"ValidateConfigResponse", string(data)}, " ")
}
