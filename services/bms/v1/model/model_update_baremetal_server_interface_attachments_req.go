package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBaremetalServerInterfaceAttachmentsReq 修改裸金属服务器所绑定的弹性网卡属性：终止时删除，即删除裸金属服务器，或者解绑网卡时是否删除该网卡
type UpdateBaremetalServerInterfaceAttachmentsReq struct {
	InterfaceAttachment *InterfaceAttachmentsReq `json:"interface_attachment"`
}

func (o UpdateBaremetalServerInterfaceAttachmentsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBaremetalServerInterfaceAttachmentsReq struct{}"
	}

	return strings.Join([]string{"UpdateBaremetalServerInterfaceAttachmentsReq", string(data)}, " ")
}
