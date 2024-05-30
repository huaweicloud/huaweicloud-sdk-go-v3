package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StandElementValueVoList 属性值列表。
type StandElementValueVoList struct {

	// 属性信息。
	Values []StandElementValueVo `json:"values"`

	// 数据标准的ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 标准所属目录，填写String类型替代Long类型。
	DirectoryId string `json:"directory_id"`

	// 目录树。
	DirectoryPath *string `json:"directory_path,omitempty"`

	// 标准行的ID，填写String类型替代Long类型。
	RowId *string `json:"row_id,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	// 是否来自公共层。
	FromPublic *bool `json:"from_public,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o StandElementValueVoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StandElementValueVoList struct{}"
	}

	return strings.Join([]string{"StandElementValueVoList", string(data)}, " ")
}
