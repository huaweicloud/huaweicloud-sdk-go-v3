package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGeipResourceQuotasResponse Response Object
type ListGeipResourceQuotasResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	Quotas *ListQuotas `json:"quotas,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGeipResourceQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeipResourceQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListGeipResourceQuotasResponse", string(data)}, " ")
}
