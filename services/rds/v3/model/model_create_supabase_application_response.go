package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSupabaseApplicationResponse Response Object
type CreateSupabaseApplicationResponse struct {

	// 订单ID（包周期时返回）。
	OrderId *string `json:"order_id,omitempty"`

	// 应用创建任务ID。
	JobId *string `json:"job_id,omitempty"`

	// Supabase应用ID。
	ApplicationId *string `json:"application_id,omitempty"`

	// 数据库实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例创建任务ID。
	InstanceJobId  *string `json:"instance_job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSupabaseApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSupabaseApplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateSupabaseApplicationResponse", string(data)}, " ")
}
