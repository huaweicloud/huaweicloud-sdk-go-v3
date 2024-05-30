package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageMemberResponse Response Object
type DeleteImageMemberResponse struct {

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageMemberResponse struct{}"
	}

	return strings.Join([]string{"DeleteImageMemberResponse", string(data)}, " ")
}
