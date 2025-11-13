package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteServerGroupMemberResponse Response Object
type BatchDeleteServerGroupMemberResponse struct {
	Status *string `json:"status,omitempty"`

	Servers        *[]BatchOperateResultResponse `json:"servers,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o BatchDeleteServerGroupMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerGroupMemberResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerGroupMemberResponse", string(data)}, " ")
}
