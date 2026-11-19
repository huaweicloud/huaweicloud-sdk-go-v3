package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyOpsWindowResponse Response Object
type ModifyOpsWindowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyOpsWindowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyOpsWindowResponse struct{}"
	}

	return strings.Join([]string{"ModifyOpsWindowResponse", string(data)}, " ")
}
