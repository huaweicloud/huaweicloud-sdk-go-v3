package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateAppGroupRequest Request Object
type DisassociateAppGroupRequest struct {

	// 服务器组ID。
	ServerGroupId string `json:"server_group_id"`
}

func (o DisassociateAppGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateAppGroupRequest struct{}"
	}

	return strings.Join([]string{"DisassociateAppGroupRequest", string(data)}, " ")
}
