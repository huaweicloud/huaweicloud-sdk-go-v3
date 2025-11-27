package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateModuleVersionResponse Response Object
type DeletePrivateModuleVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateModuleVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateModuleVersionResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateModuleVersionResponse", string(data)}, " ")
}
