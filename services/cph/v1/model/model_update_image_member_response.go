package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImageMemberResponse Response Object
type UpdateImageMemberResponse struct {

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageMemberResponse struct{}"
	}

	return strings.Join([]string{"UpdateImageMemberResponse", string(data)}, " ")
}
