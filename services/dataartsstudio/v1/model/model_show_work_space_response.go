package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkSpaceResponse Response Object
type ShowWorkSpaceResponse struct {

	// DLI脏数据OBS路径
	BadRecordLocationName *string `json:"bad_record_location_name,omitempty"`

	// 工作空间描述
	Description *string `json:"description,omitempty"`

	// 企业项目id，如果当前为公有云，且用户开启企业项目，则必选
	EpsId *string `json:"eps_id,omitempty"`

	// 作业日志OBS路径
	JobLogLocationName *string `json:"job_log_location_name,omitempty"`

	// 工作空间名称
	Name *string `json:"name,omitempty"`

	// 工作空间id
	Id *string `json:"id,omitempty"`

	// 是否为默认空间，0为私有空间，1为默认空间，2为公共空间
	IsDefault *int32 `json:"is_default,omitempty"`

	// 创建者名称
	OwnerName *string `json:"owner_name,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 当前租户所属domain id
	DomainId *string `json:"domain_id,omitempty"`

	// 当前工作空间所属实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 创建时间
	CreateTime float32 `json:"create_time,omitempty"`

	// 创建用户名称
	CreateUser *string `json:"create_user,omitempty"`

	// 当前工作空间成员数量
	MemberNum *int32 `json:"member_num,omitempty"`

	// 更新时间
	UpdateTime float32 `json:"update_time,omitempty"`

	// 更新用户名称
	UpdateUser     *string `json:"update_user,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWorkSpaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkSpaceResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkSpaceResponse", string(data)}, " ")
}
