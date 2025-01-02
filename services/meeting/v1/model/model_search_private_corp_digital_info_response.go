package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchPrivateCorpDigitalInfoResponse Response Object
type SearchPrivateCorpDigitalInfoResponse struct {

	// 结果码
	ReturnCode *int32 `json:"returnCode,omitempty"`

	// 结果描述
	ReturnDesc *string `json:"returnDesc,omitempty"`

	// 数字资产列表
	CorpDigitalInfoList *[]CorpDigitalInfo `json:"corpDigitalInfoList,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o SearchPrivateCorpDigitalInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPrivateCorpDigitalInfoResponse struct{}"
	}

	return strings.Join([]string{"SearchPrivateCorpDigitalInfoResponse", string(data)}, " ")
}
