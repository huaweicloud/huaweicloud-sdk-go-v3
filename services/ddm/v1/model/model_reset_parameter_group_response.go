package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetParameterGroupResponse Response Object
type ResetParameterGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetParameterGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetParameterGroupResponse struct{}"
	}

	return strings.Join([]string{"ResetParameterGroupResponse", string(data)}, " ")
}
