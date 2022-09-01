package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMetadataResponse struct {

	// Metadata的ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 身份提供商ID。
	IdpId *string `json:"idp_id,omitempty" xml:"idp_id"`

	// Metadata文件中的entityID字段。
	EntityId *string `json:"entity_id,omitempty" xml:"entity_id"`

	// 协议ID。
	ProtocolId *string `json:"protocol_id,omitempty" xml:"protocol_id"`

	// 用户所属账号ID。
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 账号来源，默认为空。
	XaccountType *string `json:"xaccount_type,omitempty" xml:"xaccount_type"`

	// 导入或更新Metadata文件的时间。
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	// Metadata文件的内容。
	Data           *string `json:"data,omitempty" xml:"data"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetadataResponse struct{}"
	}

	return strings.Join([]string{"ShowMetadataResponse", string(data)}, " ")
}
