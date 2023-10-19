package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAclAccountResponse Response Object
type DeleteAclAccountResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAclAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAclAccountResponse struct{}"
	}

	return strings.Join([]string{"DeleteAclAccountResponse", string(data)}, " ")
}
