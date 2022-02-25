package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateResourceRequest struct {
	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成

	XRequestId *string `json:"X-Request-Id,omitempty"`
	// 语言参数，默认为中文zh-CN, 英文为en-US

	AcceptLanguage *string `json:"Accept-Language,omitempty"`
	// 企业id

	CorpId string `json:"corp_id"`
	// 待修改的资源列表，最多支持批量处理100个

	Body *[]ModResourceDto `json:"body,omitempty"`
}

func (o UpdateResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceRequest struct{}"
	}

	return strings.Join([]string{"UpdateResourceRequest", string(data)}, " ")
}
