package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAclAccountResponse Response Object
type CreateAclAccountResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAclAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAclAccountResponse struct{}"
	}

	return strings.Join([]string{"CreateAclAccountResponse", string(data)}, " ")
}
