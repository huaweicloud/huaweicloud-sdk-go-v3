package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CancelScriptResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelScriptResponse struct{}"
	}

	return strings.Join([]string{"CancelScriptResponse", string(data)}, " ")
}
