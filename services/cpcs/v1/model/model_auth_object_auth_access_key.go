package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthObjectAuthAccessKey ak/sk信息
type AuthObjectAuthAccessKey struct {

	// ak id
	Id *string `json:"id,omitempty"`

	// sk id
	Secret *string `json:"secret,omitempty"`
}

func (o AuthObjectAuthAccessKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthObjectAuthAccessKey struct{}"
	}

	return strings.Join([]string{"AuthObjectAuthAccessKey", string(data)}, " ")
}
