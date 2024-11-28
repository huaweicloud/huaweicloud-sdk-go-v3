package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubscriptionResponse Response Object
type ShowSubscriptionResponse struct {

	// forcePurchase
	ForcePurchase *bool `json:"forcePurchase,omitempty"`

	// amount
	Amount *int32 `json:"amount,omitempty"`

	// expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// resources
	Resources *[]ShowSubscriptionResources `json:"resources,omitempty"`

	// used
	Used *int32 `json:"used,omitempty"`

	// type
	Type *int32 `json:"type,omitempty"`

	// isNewUser
	IsNewUser *bool `json:"isNewUser,omitempty"`

	// version
	Version        *string `json:"version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionResponse", string(data)}, " ")
}
