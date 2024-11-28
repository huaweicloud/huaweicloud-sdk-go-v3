package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowSubscriptionResources struct {

	// resourceId
	ResourceId *string `json:"resourceId,omitempty"`

	// resourceSpecCode
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty"`

	// resourceType
	ResourceType *string `json:"resourceType,omitempty"`

	// resourceSize
	ResourceSize *int32 `json:"resourceSize,omitempty"`

	// expireTime
	ExpireTime *string `json:"expireTime,omitempty"`

	// status
	Status *int32 `json:"status,omitempty"`
}

func (o ShowSubscriptionResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubscriptionResources struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionResources", string(data)}, " ")
}
