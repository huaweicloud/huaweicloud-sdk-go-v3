package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 归档信息
type BackupDto struct {

	// id
	Id *string `json:"id,omitempty"`

	// 归档名称
	Name *string `json:"name,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 归档数据路径集
	Paths *[]string `json:"paths,omitempty"`

	// 归档开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 归档结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 大小
	Size *int64 `json:"size,omitempty"`

	// 归档描述
	Description *string `json:"description,omitempty"`

	// 归档人员姓名
	OperatorName *string `json:"operator_name,omitempty"`
}

func (o BackupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupDto struct{}"
	}

	return strings.Join([]string{"BackupDto", string(data)}, " ")
}
