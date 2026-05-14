package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindLogicDbResponse Response Object
type UnbindLogicDbResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UnbindLogicDbResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindLogicDbResponse struct{}"
	}

	return strings.Join([]string{"UnbindLogicDbResponse", string(data)}, " ")
}
