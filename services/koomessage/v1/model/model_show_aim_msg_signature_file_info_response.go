package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAimMsgSignatureFileInfoResponse Response Object
type ShowAimMsgSignatureFileInfoResponse struct {

	// 文件ID。
	FileId *string `json:"file_id,omitempty"`

	// 文件名称。
	FileName       *string `json:"file_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAimMsgSignatureFileInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAimMsgSignatureFileInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowAimMsgSignatureFileInfoResponse", string(data)}, " ")
}
