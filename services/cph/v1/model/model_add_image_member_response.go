package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddImageMemberResponse Response Object
type AddImageMemberResponse struct {

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageMemberResponse struct{}"
	}

	return strings.Join([]string{"AddImageMemberResponse", string(data)}, " ")
}
