package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodeActivationRecordsResponse Response Object
type ShowNodeActivationRecordsResponse struct {

	// 满足条件的激活记录总数
	Count *int32 `json:"count,omitempty"`

	// 激活记录列表
	Records        *[]ActivateRecordRecords `json:"records,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowNodeActivationRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeActivationRecordsResponse struct{}"
	}

	return strings.Join([]string{"ShowNodeActivationRecordsResponse", string(data)}, " ")
}
