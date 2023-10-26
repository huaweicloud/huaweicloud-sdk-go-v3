package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHpcShareRequestBody 更新 HPC 型文件系统请求体
type UpdateHpcShareRequestBody struct {

	// 更新 HPC 型文件系统的操作类型。当前仅支持取值 config_gc_time
	Action string `json:"action"`

	// 文件系统冷数据淘汰时间，单位为小时，取值范围 [1, 100000000]。新建立的 OBS 绑定关系冷数据淘汰时间默认为 60 小时
	GcTime int32 `json:"gc_time"`
}

func (o UpdateHpcShareRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHpcShareRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHpcShareRequestBody", string(data)}, " ")
}
