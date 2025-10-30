package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeHostIgnoreStatusResponse Response Object
type ChangeHostIgnoreStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeHostIgnoreStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeHostIgnoreStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeHostIgnoreStatusResponse", string(data)}, " ")
}
