package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunActionsOnGroupResponse Response Object
type RunActionsOnGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RunActionsOnGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunActionsOnGroupResponse struct{}"
	}

	return strings.Join([]string{"RunActionsOnGroupResponse", string(data)}, " ")
}
