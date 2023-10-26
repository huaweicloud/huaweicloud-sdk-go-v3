package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReqUpdateHpcCacheData 后端挂载地址信息
type ReqUpdateHpcCacheData struct {

	// 冷数据淘汰时间。单位：小时。指定时间内线上缓存的数据如果没有被访问则会自动从缓存中删除。0表示数据不会因为时间原因自动从缓存中删除。
	GcTime int64 `json:"gc_time"`

	// 后端校验时间。单位：秒。指定时间间隔进行线上缓存文件与后端存储文件比较，存在变化则自动更新。0表示文件进行实时校验。
	CkTime int64 `json:"ck_time"`

	// 配置 nas 后端的信息, 和 obs 字段为二选一的关系
	Nas *[]ConfigNasTarget `json:"nas,omitempty"`

	// 配置 obs 后端的信息, 和 nas 字段为二选一的关系
	Obs *[]ConfigObsTarget `json:"obs,omitempty"`
}

func (o ReqUpdateHpcCacheData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqUpdateHpcCacheData struct{}"
	}

	return strings.Join([]string{"ReqUpdateHpcCacheData", string(data)}, " ")
}
