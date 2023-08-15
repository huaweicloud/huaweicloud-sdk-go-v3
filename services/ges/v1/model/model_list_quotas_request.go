package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotasRequest Request Object
type ListQuotasRequest struct {
}

func (o ListQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListQuotasRequest", string(data)}, " ")
}
