package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ErrorTransInfo binlog解析错误信息
type ErrorTransInfo struct {

	// 解析任务ID
	Id *int64 `json:"id,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 对象键
	ObjectKey *string `json:"object_key,omitempty"`

	// 解析开始位置
	BeginPosition *int64 `json:"begin_position,omitempty"`

	// 解析结束位置
	EndPosition *int64 `json:"end_position,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 表名称
	TbName *string `json:"tb_name,omitempty"`

	// 任务创建时间，单位毫秒
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ErrorTransInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorTransInfo struct{}"
	}

	return strings.Join([]string{"ErrorTransInfo", string(data)}, " ")
}
