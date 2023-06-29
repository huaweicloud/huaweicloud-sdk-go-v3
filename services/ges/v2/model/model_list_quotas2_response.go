package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotas2Response Response Object
type ListQuotas2Response struct {
	Quotas         *ListQuotasRespQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListQuotas2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotas2Response struct{}"
	}

	return strings.Join([]string{"ListQuotas2Response", string(data)}, " ")
}
