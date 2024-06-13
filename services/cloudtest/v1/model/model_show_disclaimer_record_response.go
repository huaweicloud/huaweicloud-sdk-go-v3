package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisclaimerRecordResponse Response Object
type ShowDisclaimerRecordResponse struct {

	// 实际的数据类型：单个对象，集合 或 NULL
	Value          *bool `json:"value,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowDisclaimerRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisclaimerRecordResponse struct{}"
	}

	return strings.Join([]string{"ShowDisclaimerRecordResponse", string(data)}, " ")
}
