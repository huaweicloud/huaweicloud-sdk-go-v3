package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchAddMembersV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddMembersV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddMembersV4Response struct{}"
	}

	return strings.Join([]string{"BatchAddMembersV4Response", string(data)}, " ")
}
