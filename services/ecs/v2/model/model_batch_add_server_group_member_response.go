package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddServerGroupMemberResponse Response Object
type BatchAddServerGroupMemberResponse struct {
	Status *string `json:"status,omitempty"`

	Servers        *[]BatchOperateResultResponse `json:"servers,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o BatchAddServerGroupMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddServerGroupMemberResponse struct{}"
	}

	return strings.Join([]string{"BatchAddServerGroupMemberResponse", string(data)}, " ")
}
