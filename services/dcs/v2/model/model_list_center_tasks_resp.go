package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCenterTasksResp 任务详情数组
type ListCenterTasksResp struct {

	// 后台任务ID
	Id *string `json:"id,omitempty"`

	// 后台任务名，目前支持以下取值：  EXTEND：变更规格  BindEip：开启公网访问  UnBindEip：关闭公网访问  AddReplica：添加副本  DelReplica：删除副本  AddWhitelist：设置IP白名单  UpdatePort：修改端口  RemoveIpFromDns：域名摘除IP  masterStandbySwapJob: 主备切换任务  modify：修改密码
	Name *string `json:"name,omitempty"`

	Details *DetailsBody `json:"details,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 任务相关参数
	Params *string `json:"params,omitempty"`

	// 任务状态
	Status *string `json:"status,omitempty"`

	// 任务启动时间，格式为2020-06-17T07:38:42.503Z
	CreatedAt *string `json:"created_at,omitempty"`

	// 任务结束时间，格式为2020-06-17T07:38:42.503Z
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 错误代码
	ErrorCode *string `json:"error_code,omitempty"`

	// 是否有详细任务进展，可以展开查看
	EnableShow *bool `json:"enable_show,omitempty"`

	// 任务ID
	JobId *string `json:"job_id,omitempty"`
}

func (o ListCenterTasksResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCenterTasksResp struct{}"
	}

	return strings.Join([]string{"ListCenterTasksResp", string(data)}, " ")
}
