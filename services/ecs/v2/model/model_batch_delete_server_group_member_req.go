package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteServerGroupMemberReq struct {
	ServerIds []string `json:"server_ids"`
}

func (o BatchDeleteServerGroupMemberReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerGroupMemberReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerGroupMemberReq", string(data)}, " ")
}
