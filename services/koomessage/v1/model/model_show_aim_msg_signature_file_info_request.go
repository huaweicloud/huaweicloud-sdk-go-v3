package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAimMsgSignatureFileInfoRequest Request Object
type ShowAimMsgSignatureFileInfoRequest struct {

	// 营业执照/授权委托书文件ID。
	FileId string `json:"file_id"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *string `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *string `json:"limit,omitempty"`
}

func (o ShowAimMsgSignatureFileInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAimMsgSignatureFileInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowAimMsgSignatureFileInfoRequest", string(data)}, " ")
}
