package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SrDataStoresDatastores struct {

	// 数据库版本ID，该字段不会有重复。
	Id *string `json:"id,omitempty"`

	// 数据库版本号，只返回两位数的大版本号。
	Name *string `json:"name,omitempty"`

	// 数据库内核版本号，返回完整的四位版本号。
	KernelVersion *string `json:"kernel_version,omitempty"`
}

func (o SrDataStoresDatastores) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SrDataStoresDatastores struct{}"
	}

	return strings.Join([]string{"SrDataStoresDatastores", string(data)}, " ")
}
