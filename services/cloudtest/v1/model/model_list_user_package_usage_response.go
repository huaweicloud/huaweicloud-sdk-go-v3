package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserPackageUsageResponse Response Object
type ListUserPackageUsageResponse struct {

	// 是否请求成功
	Status *string `json:"status,omitempty"`

	// 套餐用量信息列表
	Result *[]PackageUsage `json:"result,omitempty"`

	Error          *CommonResponseErrorOfApiTest `json:"error,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListUserPackageUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserPackageUsageResponse struct{}"
	}

	return strings.Join([]string{"ListUserPackageUsageResponse", string(data)}, " ")
}
