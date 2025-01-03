package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCorpDigitalInfoListResponse Response Object
type SearchCorpDigitalInfoListResponse struct {

	// 结果码
	ReturnCode *int32 `json:"returnCode,omitempty"`

	// 结果描述
	ReturnDesc *string `json:"returnDesc,omitempty"`

	// 数字资产列表
	CorpDigitalInfoList *[]CorpDigitalInfo `json:"corpDigitalInfoList,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o SearchCorpDigitalInfoListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCorpDigitalInfoListResponse struct{}"
	}

	return strings.Join([]string{"SearchCorpDigitalInfoListResponse", string(data)}, " ")
}
