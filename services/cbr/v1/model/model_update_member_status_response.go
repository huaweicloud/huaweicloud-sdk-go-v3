package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMemberStatusResponse Response Object
type UpdateMemberStatusResponse struct {
	Member         *Member `json:"member,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMemberStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateMemberStatusResponse", string(data)}, " ")
}
