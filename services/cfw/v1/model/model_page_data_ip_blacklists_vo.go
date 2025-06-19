package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageDataIpBlacklistsVo 查询出的IP黑名单列表信息，total字段表示已经导入的IP黑名单条数，如果没有导入过则total字段为0条，最大为2条。 records字段中保存了返回的IP黑名单列表信息。
type PageDataIpBlacklistsVo struct {

	// 一次查询返回的记录条数，调用接口时赋值
	Limit *int32 `json:"limit,omitempty"`

	// 本次查询结果返回后下次查询的偏移
	Offset *int32 `json:"offset,omitempty"`

	// 查询出的IP黑名单列表信息
	Records *[]IpBlacklistVo `json:"records,omitempty"`

	// 防火墙实例中IP黑名单的总数
	Total *int32 `json:"total,omitempty"`
}

func (o PageDataIpBlacklistsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageDataIpBlacklistsVo struct{}"
	}

	return strings.Join([]string{"PageDataIpBlacklistsVo", string(data)}, " ")
}
