package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeReqDetail 设备详细信息
type NodeReqDetail struct {

	// 批量注册数量。默认为单设备注册，即batch=1，如果大于1即为批量注册的最大数量。最大上限数量为100000
	Batch *int32 `json:"batch,omitempty"`

	// 设备描述，最大长度255，不允许^, ~, ＃, $, %, &, *, <, >, (, ), [, ], {, }, ', \", \\
	Description *string `json:"description,omitempty"`

	// 是否开启GPU，true表示开启，false表示不开启，默认为false
	EnableGpu *bool `json:"enable_gpu,omitempty"`

	// 是否开启NPU，true表示开启，false表示不开启，默认为false
	EnableNpu *bool `json:"enable_npu,omitempty"`

	// 子账号ID。当主账号注册设备时，可以指定将设备注册到指定的子账号下面。不填默认为该发起注册行为用户的user_id
	IamUserId *string `json:"iam_user_id,omitempty"`

	// 设备日志配置
	LogConfigs *[]LogConfig `json:"log_configs,omitempty"`

	// 设备名称，只允许中文字符、英文字母、数字、下划线、中划线，最大长度64
	Name string `json:"name"`

	// NPU类型，D310/D910。不填表示为D310类型。
	NpuType *string `json:"npu_type,omitempty"`

	// 设备标签，标签个数最多为20个
	Tags *[]TagObject `json:"tags,omitempty"`

	// 工作空间ID，不填为为主账号/子账号的默认工作空间
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 事件有效时间（单位：分钟）
	EventValidityPeriod *int32 `json:"event_validity_period,omitempty"`
}

func (o NodeReqDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeReqDetail struct{}"
	}

	return strings.Join([]string{"NodeReqDetail", string(data)}, " ")
}
