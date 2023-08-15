package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateCodeResponse Response Object
type ValidateCodeResponse struct {

	// 预验证ticket
	Ticket         *string `json:"ticket,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ValidateCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateCodeResponse struct{}"
	}

	return strings.Join([]string{"ValidateCodeResponse", string(data)}, " ")
}
