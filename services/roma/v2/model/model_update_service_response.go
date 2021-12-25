package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateServiceResponse struct {
	// 服务归属的产品模板ID

	ProductTemplateId *int32 `json:"product_template_id,omitempty"`
	// 服务归属的产品ID

	ProductId *int32 `json:"product_id,omitempty"`
	// 服务ID

	ServiceId *int32 `json:"service_id,omitempty"`
	// 服务名称，支持中文,英文大小写，数字，下划线和中划线,长度2-64

	ServiceName *string `json:"service_name,omitempty"`
	// 服务描述，长度0-200

	Description *string `json:"description,omitempty"`
	// 服务状态 0-启用 1-停用

	Status *int32 `json:"status,omitempty"`

	CreatedUser *CreatedUser `json:"created_user,omitempty"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty"`
	// 创建时间止，格式timestamp(ms)，使用UTC时区

	CreatedDatetime *int64 `json:"created_datetime,omitempty"`
	// 最后修改时间止，格式timestamp(ms)，使用UTC时区

	LastUpdatedDatetime *int64 `json:"last_updated_datetime,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o UpdateServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceResponse struct{}"
	}

	return strings.Join([]string{"UpdateServiceResponse", string(data)}, " ")
}
