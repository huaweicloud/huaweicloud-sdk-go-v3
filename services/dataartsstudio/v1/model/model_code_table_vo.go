package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeTableVo 码表结构
type CodeTableVo struct {

	// ID
	Id *int64 `json:"id,omitempty"`

	// 表名称，英文名
	NameEn string `json:"name_en"`

	// 表名称，中文名
	NameCh string `json:"name_ch"`

	// 表版本
	TbVersion *int32 `json:"tb_version,omitempty"`

	// 目录ID
	DirectoryId int64 `json:"directory_id"`

	// 目录树
	DirectoryPath *string `json:"directory_path,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	// 码表属性信息
	CodeTableFields []CodeTableFieldVo `json:"code_table_fields"`
}

func (o CodeTableVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeTableVo struct{}"
	}

	return strings.Join([]string{"CodeTableVo", string(data)}, " ")
}
