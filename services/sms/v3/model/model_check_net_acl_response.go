package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckNetAclResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckNetAclResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckNetAclResponse struct{}"
	}

	return strings.Join([]string{"CheckNetAclResponse", string(data)}, " ")
}
