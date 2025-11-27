package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecordSetsRequest Request Object
type ListRecordSetsRequest struct {
}

func (o ListRecordSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsRequest struct{}"
	}

	return strings.Join([]string{"ListRecordSetsRequest", string(data)}, " ")
}
