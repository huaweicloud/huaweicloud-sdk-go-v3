package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationGroupResponse Response Object
type CreateApplicationGroupResponse struct {
	Data           *GroupCreateResponseData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CreateApplicationGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateApplicationGroupResponse", string(data)}, " ")
}
