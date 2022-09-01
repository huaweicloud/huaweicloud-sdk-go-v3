package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBaremetalServerInterfaceAttachmentsResponse struct {

	// 裸金属服务器网卡信息列表，详情请参见表2 interfaceAttachments字段数据结构说明。
	InterfaceAttachments *[]InterfaceAttachments `json:"interfaceAttachments,omitempty" xml:"interfaceAttachments"`
	HttpStatusCode       int                     `json:"-"`
}

func (o ShowBaremetalServerInterfaceAttachmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaremetalServerInterfaceAttachmentsResponse struct{}"
	}

	return strings.Join([]string{"ShowBaremetalServerInterfaceAttachmentsResponse", string(data)}, " ")
}
