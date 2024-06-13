package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentVo 实际的数据类型：单个对象，集合 或 NULL
type AttachmentVo struct {

	// 附件Uri
	Uri *string `json:"uri,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 更新人
	Updator *string `json:"updator,omitempty"`

	// 逻辑region
	Region *string `json:"region,omitempty"`

	// 文档id
	DocId *string `json:"doc_id,omitempty"`

	// 父节点Uri
	ParentUri *string `json:"parent_uri,omitempty"`

	// 父节点类型
	ParentType *string `json:"parent_type,omitempty"`

	// 文件名
	FileName *string `json:"file_name,omitempty"`

	// 保存文件名
	StoreFileName *string `json:"store_file_name,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 文件大小
	FileSize *int32 `json:"file_size,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// 系统区分：docman或testman
	SystemType *string `json:"system_type,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 创建时间时间戳
	CreateTimeTimestamp *int64 `json:"create_time_timestamp,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 更新时间时间戳
	UpdateTimeTimestamp *int64 `json:"update_time_timestamp,omitempty"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 附件类型 0 本地上传  other 关联文档
	RelatedType *string `json:"related_type,omitempty"`
}

func (o AttachmentVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentVo struct{}"
	}

	return strings.Join([]string{"AttachmentVo", string(data)}, " ")
}
