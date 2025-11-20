package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserQuotasResponse Response Object
type ListUserQuotasResponse struct {

	// 子账户总数。
	Count *int32 `json:"count,omitempty"`

	// 子账户列表。
	Assets *[]UserQuotaInfo `json:"assets,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListUserQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListUserQuotasResponse", string(data)}, " ")
}
