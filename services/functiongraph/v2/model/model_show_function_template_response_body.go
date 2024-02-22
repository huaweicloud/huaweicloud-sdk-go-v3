package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowFunctionTemplateResponseBody struct {

	// 模板id
	Id *string `json:"id,omitempty"`

	// 模板类型
	Type *int32 `json:"type,omitempty"`

	// 模板标题
	Title *string `json:"title,omitempty"`

	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 模板描述
	Description *string `json:"description,omitempty"`

	// 模板执行运行时
	Runtime *string `json:"runtime,omitempty"`

	// 模板函数执行入口
	Handler *string `json:"handler,omitempty"`

	// 代码类型
	CodeType *string `json:"code_type,omitempty"`

	// 代码文件
	Code *string `json:"code,omitempty"`

	// 函数执行超时时间，超时函数将被强行停止，范围3～259200秒。
	Timeout *int32 `json:"timeout,omitempty"`

	// 内存大小
	MemorySize *int32 `json:"memory_size,omitempty"`

	// 触发信息列表
	TriggerMetadataList *[]TriggerMetadataList `json:"trigger_metadata_list,omitempty"`

	TempDetail *TempDetail `json:"temp_detail,omitempty"`

	// 用户数据
	UserData *string `json:"user_data,omitempty"`

	// 加密用户数据
	EncryptedUserData *string `json:"encrypted_user_data,omitempty"`

	// 模板所需依赖列表
	Dependencies *[]string `json:"dependencies,omitempty"`

	// 模板使用场景
	Scene *string `json:"scene,omitempty"`

	// 模板关联云服务
	Service *string `json:"service,omitempty"`
}

func (o ShowFunctionTemplateResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionTemplateResponseBody struct{}"
	}

	return strings.Join([]string{"ShowFunctionTemplateResponseBody", string(data)}, " ")
}
