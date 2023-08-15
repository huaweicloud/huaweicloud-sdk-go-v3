package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccountStatusDto 包含有关在组织中创建帐号的CreateAccount请求的状态。
type CreateAccountStatusDto struct {

	// 如果帐号创建成功，则为新帐号的唯一标识符（ID）。
	AccountId string `json:"account_id"`

	// 帐号名称。
	AccountName string `json:"account_name"`

	// 创建帐号和完成请求的日期和时间。
	CompletedAt *sdktime.SdkTime `json:"completed_at"`

	// 请求创建帐号的日期和时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 请求的唯一标识符（ID）。您可以从创建帐号的初始CreateAccount请求的响应中获得此值。
	Id string `json:"id"`

	// 创建帐号的异步请求的状态，in_progress：处理中，succeeded：成功，failed：失败。
	State string `json:"state"`

	// 如果请求失败，则说明失败原因。
	FailureReason *string `json:"failure_reason,omitempty"`
}

func (o CreateAccountStatusDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccountStatusDto struct{}"
	}

	return strings.Join([]string{"CreateAccountStatusDto", string(data)}, " ")
}
