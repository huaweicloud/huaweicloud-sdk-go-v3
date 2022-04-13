package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListQuotaRequest struct {
}

func (o ListQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaRequest struct{}"
	}

	return strings.Join([]string{"ListQuotaRequest", string(data)}, " ")
}
