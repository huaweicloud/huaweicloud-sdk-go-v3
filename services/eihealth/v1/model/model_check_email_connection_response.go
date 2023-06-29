package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckEmailConnectionResponse Response Object
type CheckEmailConnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckEmailConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckEmailConnectionResponse struct{}"
	}

	return strings.Join([]string{"CheckEmailConnectionResponse", string(data)}, " ")
}
