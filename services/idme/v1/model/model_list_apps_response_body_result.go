package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAppsResponseBodyResult struct {

	// 应用ID。
	Id *string `json:"id,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 删除标记。 - 0：未删除 - 1：删除
	MarkForDelete *int32 `json:"mark_for_delete,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 应用的中文名称。
	NameCn *string `json:"name_cn,omitempty"`

	// 应用的英文名称。
	NameEn *string `json:"name_en,omitempty"`

	// 应用的中文描述。
	DescCn *string `json:"desc_cn,omitempty"`

	// 应用的英文描述。
	DescEn *string `json:"desc_en,omitempty"`

	// 应用的数据库类型。
	DatabaseType *string `json:"database_type,omitempty"`

	// 运行服务的环境标识。
	Environment *string `json:"environment,omitempty"`

	// 应用责任人。
	Owners *[]string `json:"owners,omitempty"`

	// App类型。
	AppType *string `json:"app_type,omitempty"`

	// App权限控制。
	PermissionControl *string `json:"permission_control,omitempty"`
}

func (o ListAppsResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsResponseBodyResult struct{}"
	}

	return strings.Join([]string{"ListAppsResponseBodyResult", string(data)}, " ")
}
