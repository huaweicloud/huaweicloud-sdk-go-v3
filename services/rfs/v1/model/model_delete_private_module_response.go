package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateModuleResponse Response Object
type DeletePrivateModuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateModuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateModuleResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateModuleResponse", string(data)}, " ")
}
