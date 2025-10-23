package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyHbaConfResponse Response Object
type ModifyHbaConfResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyHbaConfResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyHbaConfResponse struct{}"
	}

	return strings.Join([]string{"ModifyHbaConfResponse", string(data)}, " ")
}
