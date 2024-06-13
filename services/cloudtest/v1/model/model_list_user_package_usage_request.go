package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserPackageUsageRequest Request Object
type ListUserPackageUsageRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ListUserPackageUsageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserPackageUsageRequest struct{}"
	}

	return strings.Join([]string{"ListUserPackageUsageRequest", string(data)}, " ")
}
