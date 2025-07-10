package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteScriptOrCommandReq 批量执行脚本请求体。
type ExecuteScriptOrCommandReq struct {

	// 灰度资源数量。
	GrayCount *int32 `json:"gray_count,omitempty"`

	// 资源类型，如桌面(DESKTOP)。
	ResourceType *string `json:"resource_type,omitempty"`

	// 灰度执行脚本的资源列表。
	GrayResourceIds *[]string `json:"gray_resource_ids,omitempty"`

	// 灰度失败阈值，达到阈值后停止进行下一批资源执行脚本，小于gray_count。
	GrayFailThreshold *int32 `json:"gray_fail_threshold,omitempty"`

	// 执行脚本的资源列表。
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 执行的脚本列表。
	ScriptIds *[]string `json:"script_ids,omitempty"`

	// 执行脚本的超时时间，单位分钟。
	ExecutionTimeout *int32 `json:"execution_timeout,omitempty"`
}

func (o ExecuteScriptOrCommandReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScriptOrCommandReq struct{}"
	}

	return strings.Join([]string{"ExecuteScriptOrCommandReq", string(data)}, " ")
}
