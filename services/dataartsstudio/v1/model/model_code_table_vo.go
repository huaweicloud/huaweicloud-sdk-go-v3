package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeTableVo 码表结构。
type CodeTableVo struct {

	// 码表ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 表名称，英文名。
	NameEn string `json:"name_en"`

	// 表名称，中文名。
	NameCh string `json:"name_ch"`

	// 表版本。
	TbVersion *int32 `json:"tb_version,omitempty"`

	// 目录ID，填写String类型替代Long类型。
	DirectoryId string `json:"directory_id"`

	// 目录树。
	DirectoryPath *string `json:"directory_path,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	// 码表属性信息。
	CodeTableFields []CodeTableFieldVo `json:"code_table_fields"`
}

func (o CodeTableVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeTableVo struct{}"
	}

	return strings.Join([]string{"CodeTableVo", string(data)}, " ")
}
