package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimMsgSignatureRequest Request Object
type ListAimMsgSignatureRequest struct {

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 签名ID。
	SignatureId *string `json:"signature_id,omitempty"`

	// 签名名称。
	SignatureName *string `json:"signature_name,omitempty"`

	// 签名类型。
	SignatureType *string `json:"signature_type,omitempty"`

	// 开始时间。
	BeginTime *string `json:"begin_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *string `json:"offset,omitempty"`
}

func (o ListAimMsgSignatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimMsgSignatureRequest struct{}"
	}

	return strings.Join([]string{"ListAimMsgSignatureRequest", string(data)}, " ")
}
