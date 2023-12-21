package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginInstanceRequest Request Object
type LoginInstanceRequest struct {
	Body *CommonCbhRequestBody `json:"body,omitempty"`
}

func (o LoginInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginInstanceRequest struct{}"
	}

	return strings.Join([]string{"LoginInstanceRequest", string(data)}, " ")
}
