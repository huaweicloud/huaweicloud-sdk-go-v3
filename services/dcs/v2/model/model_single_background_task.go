package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单个任务结构体
type SingleBackgroundTask struct {

	// 后台任务ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 后台任务名，目前支持以下取值：  ChangeInstanceSpec：变更规格  BindEip：开启公网访问  UnBindEip：关闭公网访问  AddReplica：添加副本  DelReplica：删除副本  AddWhitelist：设置IP白名单  UpdatePort：修改端口  RemoveIpFromDns：域名摘除IP
	Name *string `json:"name,omitempty" xml:"name"`

	Details *DetailsBody `json:"details,omitempty" xml:"details"`

	// 用户名
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 用户ID
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 任务相关参数
	Params *string `json:"params,omitempty" xml:"params"`

	// 任务状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 任务启动时间，格式为2020-06-17T07:38:42.503Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 任务结束时间，格式为2020-06-17T07:38:42.503Z
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`
}

func (o SingleBackgroundTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SingleBackgroundTask struct{}"
	}

	return strings.Join([]string{"SingleBackgroundTask", string(data)}, " ")
}
