package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCreateAccountStatusesResponse Response Object
type ListCreateAccountStatusesResponse struct {

	// 包含有关请求的详细信息的对象列表。
	CreateAccountStatuses *[]CreateAccountStatusDto `json:"create_account_statuses,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListCreateAccountStatusesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCreateAccountStatusesResponse struct{}"
	}

	return strings.Join([]string{"ListCreateAccountStatusesResponse", string(data)}, " ")
}
