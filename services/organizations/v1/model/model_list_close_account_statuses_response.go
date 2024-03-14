package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloseAccountStatusesResponse Response Object
type ListCloseAccountStatusesResponse struct {

	// 包含有关请求的详细信息的对象列表。
	CloseAccountStatuses *[]CloseAccountStatusDto `json:"close_account_statuses,omitempty"`
	HttpStatusCode       int                      `json:"-"`
}

func (o ListCloseAccountStatusesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloseAccountStatusesResponse struct{}"
	}

	return strings.Join([]string{"ListCloseAccountStatusesResponse", string(data)}, " ")
}
