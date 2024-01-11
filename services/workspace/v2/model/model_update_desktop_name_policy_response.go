package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopNamePolicyResponse Response Object
type UpdateDesktopNamePolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDesktopNamePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopNamePolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDesktopNamePolicyResponse", string(data)}, " ")
}
