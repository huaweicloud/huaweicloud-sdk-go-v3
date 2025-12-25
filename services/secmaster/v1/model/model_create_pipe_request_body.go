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

	// 数据管道名称
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
