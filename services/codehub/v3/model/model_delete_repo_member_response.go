package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRepoMemberResponse Response Object
type DeleteRepoMemberResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *Empty `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRepoMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRepoMemberResponse struct{}"
	}

	return strings.Join([]string{"DeleteRepoMemberResponse", string(data)}, " ")
}
