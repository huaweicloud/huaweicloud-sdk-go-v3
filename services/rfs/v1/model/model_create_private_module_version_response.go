package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateModuleVersionResponse Response Object
type CreatePrivateModuleVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePrivateModuleVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateModuleVersionResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivateModuleVersionResponse", string(data)}, " ")
}
