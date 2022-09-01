package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddSshKeyResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *PublicKey `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o AddSshKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSshKeyResponse struct{}"
	}

	return strings.Join([]string{"AddSshKeyResponse", string(data)}, " ")
}
