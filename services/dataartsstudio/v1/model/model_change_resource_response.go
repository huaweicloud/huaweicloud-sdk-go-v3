package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeResourceResponse Response Object
type ChangeResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeResourceResponse struct{}"
	}

	return strings.Join([]string{"ChangeResourceResponse", string(data)}, " ")
}
