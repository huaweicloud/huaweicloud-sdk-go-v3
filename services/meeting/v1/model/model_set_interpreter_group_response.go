package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetInterpreterGroupResponse Response Object
type SetInterpreterGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetInterpreterGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetInterpreterGroupResponse struct{}"
	}

	return strings.Join([]string{"SetInterpreterGroupResponse", string(data)}, " ")
}
