package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CancelFactoryPackagesRequestBody struct {

	// 发布包id列表信息
	PackageIds *[]string `json:"package_ids,omitempty"`
}

func (o CancelFactoryPackagesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelFactoryPackagesRequestBody struct{}"
	}

	return strings.Join([]string{"CancelFactoryPackagesRequestBody", string(data)}, " ")
}
