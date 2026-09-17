package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantIssueListResponse Response Object
type ShowTenantIssueListResponse struct {

	// **参数解释**： 返回状态。 **取值范围**： success：响应成功。 error：响应失败
	Status *string `json:"status,omitempty"`

	// **参数解释**： 信息。 **取值范围**： 不涉及
	Message *string `json:"message,omitempty"`

	// **参数解释**： 查询结果。
	Result         *[]IssueDetailsResponse `json:"result,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowTenantIssueListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantIssueListResponse struct{}"
	}

	return strings.Join([]string{"ShowTenantIssueListResponse", string(data)}, " ")
}
