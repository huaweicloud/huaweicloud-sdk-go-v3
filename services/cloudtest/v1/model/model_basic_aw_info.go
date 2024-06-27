package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicAwInfo struct {

	// aw代码, hasCode为1时有效
	AwCode *string `json:"aw_code,omitempty"`

	// 模板描述
	AwDescription *string `json:"aw_description,omitempty"`

	// 组合aw包含的aw实例信息，awType为2时该字段有效
	AwInsList *[]AwInstanceRes `json:"aw_ins_list,omitempty"`

	// aw标记 0-normal；1-new；2-update 3-delete
	AwMark *int32 `json:"aw_mark,omitempty"`

	// 接口的operationId
	AwOperationid *string `json:"aw_operationid,omitempty"`

	// 接口的tags
	AwTags *string `json:"aw_tags,omitempty"`

	// AW类型1-普通aw,2-组合aw,3-导入aw
	AwType *int32 `json:"aw_type,omitempty"`

	// awOperationId_awTags的拼接
	AwUniqueid *string `json:"aw_uniqueid,omitempty"`

	// 顺序
	ByOrder *int32 `json:"by_order,omitempty"`

	// 目录层级
	CataType *int32 `json:"cata_type,omitempty"`

	// 新增childList以支持嵌套层级结构
	ChildList *[]BasicAwInfo `json:"child_list,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 时间戳字段
	CreateTimeStamp *int64 `json:"create_time_stamp,omitempty"`

	// 时间字符串
	CreateTimeString *string `json:"create_time_string,omitempty"`

	// 创建人
	CreateUser *string `json:"create_user,omitempty"`

	// 创建人id
	CreateUserId *string `json:"create_user_id,omitempty"`

	// aw库的文件列表
	CustomAwLibs *[]UploadFileRes `json:"custom_aw_libs,omitempty"`

	// 更新时间
	DeleteTime *string `json:"delete_time,omitempty"`

	// 删除人
	DeleteUser *string `json:"delete_user,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 默认检查点List
	DftCheckPointList *[]CheckPoint `json:"dft_check_point_list,omitempty"`

	// 默认请求头参数对象
	DftCustomHeader *[]AwParam `json:"dft_custom_header,omitempty"`

	// 重试间隔时间 (ms) 为空表示不等待
	DftRetryInterval *string `json:"dft_retry_interval,omitempty"`

	// 重试次数
	DftRetryTimes *string `json:"dft_retry_times,omitempty"`

	// 默认变量信息List
	DftVariableList *[]AwVariable `json:"dft_variable_list,omitempty"`

	ExtraInfo *JsonNode `json:"extra_info,omitempty"`

	// 组名
	GroupName *string `json:"group_name,omitempty"`

	// 是否自带代码 0-不自带代码，该aw依赖source字段中的源；1-自带代码
	HasCode *int32 `json:"has_code,omitempty"`

	// id
	Id *string `json:"id,omitempty"`

	// 导入的包
	ImportPackage *[]string `json:"import_package,omitempty"`

	// 接口的x-extend
	InterfaceLabel *string `json:"interface_label,omitempty"`

	// 是否是当前工程的收藏aw，该字段不存数据库，每次获取时实时判断
	IsFavorite *int32 `json:"is_favorite,omitempty"`

	// 判断是否为文件夹的标识
	IsFolder *string `json:"is_folder,omitempty"`

	// 关键字局部变量
	KeywordVariableValue *[]AwVariable `json:"keyword_variable_value,omitempty"`

	// 方法
	Method *string `json:"method,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// aw在页面上显示的名字
	NameView *string `json:"nameView,omitempty"`

	// 源工程信息，如果该aw是从其他工程过来的(继承或者copy to local)
	OriginProject *string `json:"origin_project,omitempty"`

	// 组合aw的输出参数信息，该字段不存数据库，实时获取
	OutputParamList *[]AwVariable `json:"output_param_list,omitempty"`

	// 当前页数
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页条数
	PageSize *int32 `json:"page_size,omitempty"`

	// 参数类型和参数默认值对应List
	ParamTypeAndDftValue *[]AwParam `json:"param_type_and_dft_value,omitempty"`

	// aw目录父编号
	ParentId *string `json:"parent_id,omitempty"`

	// aw目录父编号
	ProjectId *string `json:"project_id,omitempty"`

	// 协议类型 (http/dsf/other)
	ProtocolType *string `json:"protocol_type,omitempty"`

	PublicAwLib *PublicAwLibRes `json:"public_aw_lib,omitempty"`

	// 所属公共aw库Id
	PublicAwLibId *string `json:"public_aw_lib_id,omitempty"`

	// 区域名称
	Region *string `json:"region,omitempty"`

	// 返回类型
	ReturnType *string `json:"return_type,omitempty"`

	// root id
	RootId *string `json:"root_id,omitempty"`

	// 来源
	Source *string `json:"source,omitempty"`

	// 特殊AW类型 0-common,10-header,1-beforeHeader
	SpecialType *int32 `json:"special_type,omitempty"`

	// 关键字管理的用例编号
	TmssCaseNumber *string `json:"tmssCaseNumber,omitempty"`

	// 关键字aw管理的用例ID
	TmssCaseId *string `json:"tmss_case_id,omitempty"`

	// 总页数
	TotalPage *int32 `json:"total_page,omitempty"`

	// 总条数
	TotalSize *int32 `json:"total_size,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 更新时间戳
	UpdateTimeStamp *int64 `json:"update_time_stamp,omitempty"`

	// 更新字符串
	UpdateTimeString *string `json:"update_time_string,omitempty"`

	// 更新人
	UpdateUser *string `json:"update_user,omitempty"`

	// 警告信息
	WarningMsg *string `json:"warningMsg,omitempty"`

	// 所属yaml文件名称,该字段不存库，用来记录从哪个yaml文件导入
	YamlName *string `json:"yamlName,omitempty"`
}

func (o BasicAwInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicAwInfo struct{}"
	}

	return strings.Join([]string{"BasicAwInfo", string(data)}, " ")
}
