package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchAddServerGroupMemberReq struct {
	ServerIds *[]string `json:"server_ids,omitempty"`
}

func (o BatchAddServerGroupMemberReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddServerGroupMemberReq struct{}"
	}

	return strings.Join([]string{"BatchAddServerGroupMemberReq", string(data)}, " ")
}
