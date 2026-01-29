package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePipeRequestBody struct {

	// 工作空间ID
	DataspaceId string `json:"dataspace_id"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 索引字段映射；每个key对象承载一个字段的信息；存在多个key对象，key可变，表示字段名称；可嵌套
	Mapping map[string]KeyIndex `json:"mapping,omitempty"`

	// 数据管道名称；必须以英文字母开头，且只能包含小写英文字母、数字和'_'，且'_'不能在结尾，也不能连续出现。 不能以系统预留的前缀isap_、csb_、secmaster_、sec_、s_sec_、i_sec_、l_sec_、security_开头
	PipeName string `json:"pipe_name"`

	// 数据管道分区个数；默认创建1个，最大支持创建64个分区
	Shards int32 `json:"shards"`

	// 数据的保存时间，单位为天；默认30天，取值范围为7~180. 配置180天需提工单申请
	StoragePeriod int32 `json:"storage_period"`

	// 时间戳字段
	TimestampField *string `json:"timestamp_field,omitempty"`
}

func (o CreatePipeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipeRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePipeRequestBody", string(data)}, " ")
}
