package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UploadFileRes struct {

	// 所属的AWInstance的ID
	AwInsId *string `json:"awInsId,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建时间戳
	CreateTimeStamp *int64 `json:"create_time_stamp,omitempty"`

	// 创建时间字符串
	CreateTimeString *string `json:"create_time_string,omitempty"`

	// 创建人
	CreateUser *string `json:"create_user,omitempty"`

	// 当前大小
	CurrentSize *int64 `json:"current_size,omitempty"`

	// 文件路径
	FilePath *string `json:"filePath,omitempty"`

	// id
	Id *string `json:"id,omitempty"`

	// 文件在系统中的名字
	Name *string `json:"name,omitempty"`

	// 文件的原名
	OriginName *string `json:"originName,omitempty"`

	// 项目ID
	ProjectId *string `json:"projectId,omitempty"`

	// 区域名称
	Region *string `json:"region,omitempty"`

	// 测试用例的唯一标识符
	TestcaseId *string `json:"testcase_id,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 更新时间戳
	UpdateTimeStamp *int64 `json:"update_time_stamp,omitempty"`

	// 更新时间字符串
	UpdateTimeString *string `json:"update_time_string,omitempty"`

	// 更新人
	UpdateUser *string `json:"update_user,omitempty"`
}

func (o UploadFileRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFileRes struct{}"
	}

	return strings.Join([]string{"UploadFileRes", string(data)}, " ")
}
