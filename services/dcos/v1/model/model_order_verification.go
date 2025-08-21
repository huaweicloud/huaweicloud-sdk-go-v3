package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderVerification 验收信息
type OrderVerification struct {

	// 服务情况，正常或异常
	Result *string `json:"result,omitempty"`

	Details *OrderServiceDetails `json:"details,omitempty"`

	// 附件
	Attachments *[]UploadFileInfo `json:"attachments,omitempty"`

	// 是否符合预期
	MeetExpectation *bool `json:"meet_expectation,omitempty"`

	// 客户验收意见说明
	Comments *string `json:"comments,omitempty"`
}

func (o OrderVerification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderVerification struct{}"
	}

	return strings.Join([]string{"OrderVerification", string(data)}, " ")
}
