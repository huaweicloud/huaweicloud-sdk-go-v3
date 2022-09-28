package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
