package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCustomerselfResourceRecordDetailsRequest struct {
	Body *QueryResRecordsDetailReq `json:"body,omitempty"`
}

func (o ListCustomerselfResourceRecordDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerselfResourceRecordDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerselfResourceRecordDetailsRequest", string(data)}, " ")
}
