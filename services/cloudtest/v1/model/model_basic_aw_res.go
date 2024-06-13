package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicAwRes struct {
	AwCode *string `json:"aw_code,omitempty"`

	AwDescription *string `json:"aw_description,omitempty"`

	AwMark *int32 `json:"aw_mark,omitempty"`

	AwOperationid *string `json:"aw_operationid,omitempty"`

	AwTags *string `json:"aw_tags,omitempty"`

	AwType *int32 `json:"aw_type,omitempty"`

	AwUniqueid *string `json:"aw_uniqueid,omitempty"`

	ByOrder *int32 `json:"by_order,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	CreateTimeStamp *int64 `json:"create_time_stamp,omitempty"`

	CreateTimeString *string `json:"create_time_string,omitempty"`

	// 创建人
	CreateUser *string `json:"create_user,omitempty"`

	CreateUserId *string `json:"create_user_id,omitempty"`

	// 更新时间
	DeleteTime *string `json:"delete_time,omitempty"`

	// 删除人
	DeleteUser *string `json:"delete_user,omitempty"`

	Description *string `json:"description,omitempty"`

	DftCheckPointList *[]interface{} `json:"dft_check_point_list,omitempty"`

	DftCustomHeader *[]interface{} `json:"dft_custom_header,omitempty"`

	DftRetryInterval *string `json:"dft_retry_interval,omitempty"`

	DftRetryTimes *string `json:"dft_retry_times,omitempty"`

	DftVariableList *[]interface{} `json:"dft_variable_list,omitempty"`

	ExtraInfo *interface{} `json:"extra_info,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	HasCode *int32 `json:"has_code,omitempty"`

	// id
	Id *string `json:"id,omitempty"`

	ImportPackage *[]string `json:"import_package,omitempty"`

	InterfaceLabel *string `json:"interface_label,omitempty"`

	IsFavorite *int32 `json:"is_favorite,omitempty"`

	Method *string `json:"method,omitempty"`

	Name *string `json:"name,omitempty"`

	NameView *string `json:"nameView,omitempty"`

	OriginProject *string `json:"origin_project,omitempty"`

	ParamTypeAndDftValue *[]interface{} `json:"param_type_and_dft_value,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	ProtocolType *string `json:"protocol_type,omitempty"`

	PublicAwLib *interface{} `json:"public_aw_lib,omitempty"`

	PublicAwLibId *string `json:"public_aw_lib_id,omitempty"`

	Region *string `json:"region,omitempty"`

	ReturnType *string `json:"return_type,omitempty"`

	RootId *string `json:"root_id,omitempty"`

	Source *string `json:"source,omitempty"`

	SpecialType *int32 `json:"special_type,omitempty"`

	TmssCaseNumber *string `json:"tmssCaseNumber,omitempty"`

	TmssCaseId *string `json:"tmss_case_id,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	UpdateTimeStamp *int64 `json:"update_time_stamp,omitempty"`

	UpdateTimeString *string `json:"update_time_string,omitempty"`

	// 更新人
	UpdateUser *string `json:"update_user,omitempty"`

	WarningMsg *string `json:"warningMsg,omitempty"`

	YamlName *string `json:"yamlName,omitempty"`
}

func (o BasicAwRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicAwRes struct{}"
	}

	return strings.Join([]string{"BasicAwRes", string(data)}, " ")
}
