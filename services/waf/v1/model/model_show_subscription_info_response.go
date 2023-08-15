package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubscriptionInfoResponse Response Object
type ShowSubscriptionInfoResponse struct {

	// 云模式版本   - -2：已冻结   - -1：未订购   - 2：标准版   - 3：专业版   - 4：铂金版   - 7:入门版   - 22：按需版本
	Type *int32 `json:"type,omitempty"`

	// 资源列表
	Resources *[]ResourceResponse `json:"resources,omitempty"`

	// 是否为新用户
	IsNewUser *bool `json:"isNewUser,omitempty"`

	Premium        *Premium `json:"premium,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowSubscriptionInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubscriptionInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionInfoResponse", string(data)}, " ")
}
