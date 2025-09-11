package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceQuotaResponse Response Object
type ShowInstanceQuotaResponse struct {

	// CPU个数
	Cpu *int64 `json:"cpu,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 实例配额
	Quota *int64 `json:"quota,omitempty"`

	// 内存大小MB
	Ram            *int64 `json:"ram,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceQuotaResponse", string(data)}, " ")
}
