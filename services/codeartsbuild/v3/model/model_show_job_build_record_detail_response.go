package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobBuildRecordDetailResponse Response Object
type ShowJobBuildRecordDetailResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	Result         *RecordInfo2Result `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowJobBuildRecordDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobBuildRecordDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowJobBuildRecordDetailResponse", string(data)}, " ")
}
