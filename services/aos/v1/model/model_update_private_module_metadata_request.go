package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivateModuleMetadataRequest Request Object
type UpdatePrivateModuleMetadataRequest struct {

	// 用户指定的，对于此请求的唯一Id，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 私有模块（private-module）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	ModuleName string `json:"module_name"`

	Body *UpdatePrivateModuleMetadataRequestBody `json:"body,omitempty"`
}

func (o UpdatePrivateModuleMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateModuleMetadataRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateModuleMetadataRequest", string(data)}, " ")
}
