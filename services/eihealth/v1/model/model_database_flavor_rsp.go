package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseFlavorRsp struct {

	// 规格编号
	Code string `json:"code"`

	// 规格名称
	Name string `json:"name"`

	// 核数
	Cpu int32 `json:"cpu"`

	// 内存
	Ram int32 `json:"ram"`

	// 最大连接数
	MaxConnections int32 `json:"max_connections"`

	// 存储空间
	DiskSpace int32 `json:"disk_space"`

	// 是否售罄
	SoldOut bool `json:"sold_out"`
}

func (o DatabaseFlavorRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseFlavorRsp struct{}"
	}

	return strings.Join([]string{"DatabaseFlavorRsp", string(data)}, " ")
}
