package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloseAccountStatusDto 包含有关在组织中关闭账号的CloseAccount请求的状态。
type CloseAccountStatusDto struct {

	// 账号的唯一标识符（ID）。
	AccountId string `json:"account_id"`

	// 请求关闭账号的日期和时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 请求关闭账号状态更新的日期和时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 组织的唯一标识符（ID）。
	OrganizationId string `json:"organization_id"`

	// 关闭账号的状态，pending_closure：关闭中，suspended：已关闭
	State string `json:"state"`

	// 如果请求失败，则说明失败原因。
	FailureReason *string `json:"failure_reason,omitempty"`
}

func (o CloseAccountStatusDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseAccountStatusDto struct{}"
	}

	return strings.Join([]string{"CloseAccountStatusDto", string(data)}, " ")
}
