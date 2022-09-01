package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteRepoMemberResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *Empty `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRepoMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRepoMemberResponse struct{}"
	}

	return strings.Join([]string{"DeleteRepoMemberResponse", string(data)}, " ")
}
