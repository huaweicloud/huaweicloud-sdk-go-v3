package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateDesktopPoolResizeResponse Response Object
type EstimateDesktopPoolResizeResponse struct {

	// 是否为升配
	IsUpgrade *bool `json:"is_upgrade,omitempty"`

	BatchInquiryRspWhenUpgrade *PeriodBatchUpChangeResourceRsp `json:"batch_inquiry_rsp_when_upgrade,omitempty"`

	BatchInquiryRspWhenDowngrade *BatchInquiryChangeRsp `json:"batch_inquiry_rsp_when_downgrade,omitempty"`
	HttpStatusCode               int                    `json:"-"`
}

func (o EstimateDesktopPoolResizeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateDesktopPoolResizeResponse struct{}"
	}

	return strings.Join([]string{"EstimateDesktopPoolResizeResponse", string(data)}, " ")
}
