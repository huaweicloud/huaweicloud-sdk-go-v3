package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelResourcesSubscriptionResponse Response Object
type CancelResourcesSubscriptionResponse struct {

	// 客户退订订单ID的列表信息。
	OrderIds *[]string `json:"order_ids,omitempty"`

	// 失败的资源信息列表。有退订失败的资源时，该字段才有值。具体请参见表FailResourceInfo。 该字段为预留字段。
	FailResourceInfos *[]FailResourceInfo `json:"fail_resource_infos,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o CancelResourcesSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelResourcesSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"CancelResourcesSubscriptionResponse", string(data)}, " ")
}
