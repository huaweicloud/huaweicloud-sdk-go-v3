package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecordDetailResponse Response Object
type ShowRecordDetailResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	Result         *RecordInfo2Result `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowRecordDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowRecordDetailResponse", string(data)}, " ")
}
