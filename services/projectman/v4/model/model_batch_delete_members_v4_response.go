package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteMembersV4Response Response Object
type BatchDeleteMembersV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteMembersV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersV4Response struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersV4Response", string(data)}, " ")
}
