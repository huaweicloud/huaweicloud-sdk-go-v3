package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JoinGroupResponse Response Object
type JoinGroupResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o JoinGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JoinGroupResponse struct{}"
	}

	return strings.Join([]string{"JoinGroupResponse", string(data)}, " ")
}
