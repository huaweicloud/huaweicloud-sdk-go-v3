package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListErrorServersRequest struct {

	// 每一页记录的错误数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移量
	Offset int32 `json:"offset" xml:"offset"`

	// 需要查询的迁移项目id，添加此字段将只查询对应id下的迁移任务报错信息
	Migproject *string `json:"migproject,omitempty" xml:"migproject"`

	// 需要查询的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`
}

func (o ListErrorServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListErrorServersRequest struct{}"
	}

	return strings.Join([]string{"ListErrorServersRequest", string(data)}, " ")
}
