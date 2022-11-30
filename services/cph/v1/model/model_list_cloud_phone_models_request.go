package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCloudPhoneModelsRequest struct {

	// 规格状态 - 0：下线状态 - 1：正常使用 不传该参数表示查询所有状态的规格。
	Status *int32 `json:"status,omitempty"`
}

func (o ListCloudPhoneModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneModelsRequest struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneModelsRequest", string(data)}, " ")
}
