package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeclineSecurityApplicationResponse Response Object
type DeclineSecurityApplicationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeclineSecurityApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeclineSecurityApplicationResponse struct{}"
	}

	return strings.Join([]string{"DeclineSecurityApplicationResponse", string(data)}, " ")
}
