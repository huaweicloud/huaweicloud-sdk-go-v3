package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HbaseModifySettingV2Req 修改配置v2接口请求参数
type HbaseModifySettingV2Req struct {

	// 是否重启
	IsReboot string `json:"is_reboot"`

	// 参见HBaseModifySettingV2结构说明
	HbaseModifySettings []HbaseModifySettingV2 `json:"hbase_modify_settings"`
}

func (o HbaseModifySettingV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HbaseModifySettingV2Req struct{}"
	}

	return strings.Join([]string{"HbaseModifySettingV2Req", string(data)}, " ")
}
