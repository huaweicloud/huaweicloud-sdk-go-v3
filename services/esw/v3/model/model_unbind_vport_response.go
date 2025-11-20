package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindVportResponse Response Object
type UnbindVportResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UnbindVportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindVportResponse struct{}"
	}

	return strings.Join([]string{"UnbindVportResponse", string(data)}, " ")
}
