package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StandElementValueVoList 属性值列表
type StandElementValueVoList struct {

	// 属性信息
	Values []StandElementValueVo `json:"values"`

	// ID
	Id *int64 `json:"id,omitempty"`

	// 标准所属目录
	DirectoryId int64 `json:"directory_id"`

	// 目录树
	DirectoryPath *string `json:"directory_path,omitempty"`

	// 标准行的id
	RowId *int64 `json:"row_id,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o StandElementValueVoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StandElementValueVoList struct{}"
	}

	return strings.Join([]string{"StandElementValueVoList", string(data)}, " ")
}
