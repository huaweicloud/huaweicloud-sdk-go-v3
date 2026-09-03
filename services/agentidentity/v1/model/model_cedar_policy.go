package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CedarPolicy struct {

	// The Cedar policy statement that defines the authorization logic. This statement follows Cedar syntax and specifies principals, actions, resources, and conditions for fine-grained access control.
	Statement string `json:"statement"`
}

func (o CedarPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CedarPolicy struct{}"
	}

	return strings.Join([]string{"CedarPolicy", string(data)}, " ")
}
