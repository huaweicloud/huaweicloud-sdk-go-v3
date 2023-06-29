package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProjectModuleResponse Response Object
type DeleteProjectModuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProjectModuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProjectModuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteProjectModuleResponse", string(data)}, " ")
}
