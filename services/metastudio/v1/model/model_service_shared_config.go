package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceSharedConfig 业务共享配置
type ServiceSharedConfig struct {

	// 开启共享配置
	Enable *bool `json:"enable,omitempty"`

	// 可选的共享租户列表
	OptionalProjectIds *[]string `json:"optional_project_ids,omitempty"`

	// 共享免审核
	NoNeedReview *bool `json:"no_need_review,omitempty"`
}

func (o ServiceSharedConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceSharedConfig struct{}"
	}

	return strings.Join([]string{"ServiceSharedConfig", string(data)}, " ")
}
