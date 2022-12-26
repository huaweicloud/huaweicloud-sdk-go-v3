package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHistorySqlRequest struct {
}

func (o ListHistorySqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistorySqlRequest struct{}"
	}

	return strings.Join([]string{"ListHistorySqlRequest", string(data)}, " ")
}
