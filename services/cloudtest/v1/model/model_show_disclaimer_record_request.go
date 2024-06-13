package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisclaimerRecordRequest Request Object
type ShowDisclaimerRecordRequest struct {

	// 类型
	Type string `json:"type"`
}

func (o ShowDisclaimerRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisclaimerRecordRequest struct{}"
	}

	return strings.Join([]string{"ShowDisclaimerRecordRequest", string(data)}, " ")
}
