package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDisableMembersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDisableMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisableMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchDisableMembersResponse", string(data)}, " ")
}
