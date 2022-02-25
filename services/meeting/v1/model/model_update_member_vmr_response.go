package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateMemberVmrResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMemberVmrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberVmrResponse struct{}"
	}

	return strings.Join([]string{"UpdateMemberVmrResponse", string(data)}, " ")
}
