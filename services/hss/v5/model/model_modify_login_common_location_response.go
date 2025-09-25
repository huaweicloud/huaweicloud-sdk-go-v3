package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyLoginCommonLocationResponse Response Object
type ModifyLoginCommonLocationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyLoginCommonLocationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyLoginCommonLocationResponse struct{}"
	}

	return strings.Join([]string{"ModifyLoginCommonLocationResponse", string(data)}, " ")
}
