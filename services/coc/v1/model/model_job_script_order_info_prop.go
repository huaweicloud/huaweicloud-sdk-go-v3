package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobScriptOrderInfoProp 基本信息
type JobScriptOrderInfoProp struct {

	// 脚本uuid
	ScriptUuid *string `json:"script_uuid,omitempty"`

	// 脚本名称
	ScriptName *string `json:"script_name,omitempty"`

	// 脚本版本uuid
	ScriptVersionUuid *int64 `json:"script_version_uuid,omitempty"`

	// 脚本版本名
	ScriptVersionName *string `json:"script_version_name,omitempty"`

	// 当前执行批次index
	CurrentExecuteBatchIndex *int32 `json:"current_execute_batch_index,omitempty"`

	ExecuteParam *ScriptExecuteParam `json:"execute_param,omitempty"`
}

func (o JobScriptOrderInfoProp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScriptOrderInfoProp struct{}"
	}

	return strings.Join([]string{"JobScriptOrderInfoProp", string(data)}, " ")
}
