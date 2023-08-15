package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteMemberResponse Response Object
type BatchDeleteMemberResponse struct {
	Body           *[]BatchDeleteMemberRsp `json:"body,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchDeleteMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMemberResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteMemberResponse", string(data)}, " ")
}
