package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListOsVersionsRequest struct {

	// OS的标签。 根据标签值可以过滤查询指定特性的OS信息。 取值范围： bms：表示该镜像支持BMS的os_version列表。 uefi：支持UEFI启动方式的os_version列表。 arm：显示基于arm架构的os_version列表。 x86：显示基于x86架构的os_version列表。不带tag查询条件则默认查询当前region支持的所有的OS列表。
	Tag *string `json:"tag,omitempty"`
}

func (o ListOsVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOsVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListOsVersionsRequest", string(data)}, " ")
}
