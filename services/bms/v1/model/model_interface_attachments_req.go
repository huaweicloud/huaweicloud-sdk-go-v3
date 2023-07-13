package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InterfaceAttachmentsReq 删除裸金属服务器或解绑网卡时是否删除该网卡
type InterfaceAttachmentsReq struct {
	DeleteOnTermination bool `json:"delete_on_termination"`
}

func (o InterfaceAttachmentsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterfaceAttachmentsReq struct{}"
	}

	return strings.Join([]string{"InterfaceAttachmentsReq", string(data)}, " ")
}
