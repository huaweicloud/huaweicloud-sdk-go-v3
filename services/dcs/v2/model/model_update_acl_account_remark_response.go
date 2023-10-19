package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAclAccountRemarkResponse Response Object
type UpdateAclAccountRemarkResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAclAccountRemarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclAccountRemarkResponse struct{}"
	}

	return strings.Join([]string{"UpdateAclAccountRemarkResponse", string(data)}, " ")
}
