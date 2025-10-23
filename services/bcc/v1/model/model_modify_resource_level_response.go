package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyResourceLevelResponse Response Object
type ModifyResourceLevelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyResourceLevelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyResourceLevelResponse struct{}"
	}

	return strings.Join([]string{"ModifyResourceLevelResponse", string(data)}, " ")
}
