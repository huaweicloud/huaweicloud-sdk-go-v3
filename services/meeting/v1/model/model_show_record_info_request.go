package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecordInfoRequest Request Object
type ShowRecordInfoRequest struct {
	Body *RecordInfoReq `json:"body,omitempty"`
}

func (o ShowRecordInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordInfoRequest", string(data)}, " ")
}
