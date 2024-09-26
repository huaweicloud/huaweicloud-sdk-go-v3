package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecordInfoResponse Response Object
type ShowRecordInfoResponse struct {

	// 结果码
	Code *string `json:"code,omitempty"`

	// 结果描述
	Message *string `json:"message,omitempty"`

	Data           *RecordInfoDo `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowRecordInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowRecordInfoResponse", string(data)}, " ")
}
