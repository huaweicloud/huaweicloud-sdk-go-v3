package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRecordRulesResponse struct {

	// 录制规则列表
	Rules *[]RecordRule `json:"rules,omitempty" xml:"rules"`

	XRequestId     *string `json:"X-request-Id,omitempty" xml:"X-request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRecordRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordRulesResponse struct{}"
	}

	return strings.Join([]string{"ListRecordRulesResponse", string(data)}, " ")
}
