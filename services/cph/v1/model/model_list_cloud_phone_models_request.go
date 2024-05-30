package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneModelsRequest Request Object
type ListCloudPhoneModelsRequest struct {

	// 规格状态 - 0：下线状态 - 1：正常使用 不传该参数表示查询所有状态的规格。
	Status *int32 `json:"status,omitempty"`

	// 偏移量为一个大于等于0整数，表示查询该偏移量后面的所有的资源数，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页返回的资源个数。取值范围：1~100（默认值为100），一般设置为10、20、50。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCloudPhoneModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneModelsRequest struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneModelsRequest", string(data)}, " ")
}
