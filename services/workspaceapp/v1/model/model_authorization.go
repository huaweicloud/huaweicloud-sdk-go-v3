package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Authorization 用户授权信息
type Authorization struct {

	// 用户(组)
	Account *string `json:"account,omitempty"`

	// 应用ID (按照组授权时,该字段为空)
	AppId *string `json:"app_id,omitempty"`

	// 应用名称 (按照组授权时,该字段为空)
	AppName *string `json:"app_name,omitempty"`

	// 应用组ID
	AppGroupId *string `json:"app_group_id,omitempty"`

	// 应用组名称
	AppGroupName *string `json:"app_group_name,omitempty"`

	AuthorizationType *AuthorizationTypeEnum `json:"authorization_type,omitempty"`

	AccountType *AccountTypeEnum `json:"account_type,omitempty"`

	// 域名城
	Domain *string `json:"domain,omitempty"`

	// 发布时间
	CreateAt *sdktime.SdkTime `json:"create_at,omitempty"`
}

func (o Authorization) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Authorization struct{}"
	}

	return strings.Join([]string{"Authorization", string(data)}, " ")
}
