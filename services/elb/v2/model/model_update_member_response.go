package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateMemberResponse struct {
	Member         *MemberResp `json:"member,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberResponse struct{}"
	}

	return strings.Join([]string{"UpdateMemberResponse", string(data)}, " ")
}
