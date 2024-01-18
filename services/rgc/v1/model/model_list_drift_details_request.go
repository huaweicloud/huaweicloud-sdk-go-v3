package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDriftDetailsRequest Request Object
type ListDriftDetailsRequest struct {
}

func (o ListDriftDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDriftDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListDriftDetailsRequest", string(data)}, " ")
}
