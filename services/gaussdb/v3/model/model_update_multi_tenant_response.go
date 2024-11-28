package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMultiTenantResponse Response Object
type UpdateMultiTenantResponse struct {

	// 开启/关闭多租特性ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMultiTenantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMultiTenantResponse struct{}"
	}

	return strings.Join([]string{"UpdateMultiTenantResponse", string(data)}, " ")
}
