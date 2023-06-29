package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteModuleResponse Response Object
type DeleteModuleResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteModuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteModuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteModuleResponse", string(data)}, " ")
}
