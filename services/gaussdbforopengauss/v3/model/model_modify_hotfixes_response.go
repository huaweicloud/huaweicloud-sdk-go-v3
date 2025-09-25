package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyHotfixesResponse Response Object
type ModifyHotfixesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyHotfixesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyHotfixesResponse struct{}"
	}

	return strings.Join([]string{"ModifyHotfixesResponse", string(data)}, " ")
}
