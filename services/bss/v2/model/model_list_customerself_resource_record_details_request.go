package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerselfResourceRecordDetailsRequest Request Object
type ListCustomerselfResourceRecordDetailsRequest struct {

	// 大陆站默认中文，枚举：zh_cn：中文 en_us：英文
	XLanguage *string `json:"X-Language,omitempty"`

	Body *QueryResRecordsDetailReq `json:"body,omitempty"`
}

func (o ListCustomerselfResourceRecordDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerselfResourceRecordDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerselfResourceRecordDetailsRequest", string(data)}, " ")
}
