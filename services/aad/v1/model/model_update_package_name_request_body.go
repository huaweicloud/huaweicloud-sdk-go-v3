package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePackageNameRequestBody 更新实例名的请求体
type UpdatePackageNameRequestBody struct {

	// 名字
	Name string `json:"name"`
}

func (o UpdatePackageNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePackageNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePackageNameRequestBody", string(data)}, " ")
}
