package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCollectorConnectionResponse Response Object
type ShowCollectorConnectionResponse struct {

	// UUID
	ConnectionId *string `json:"connection_id,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 相关描述信息
	Fields *[]ConnectionModuleFieldShowVo `json:"fields,omitempty"`

	// UUID
	ModuleId *string `json:"module_id,omitempty"`

	// 所属租户名称
	Name *string `json:"name,omitempty"`

	// UUID
	TemplateId *string `json:"template_id,omitempty"`

	// 相关标题
	Title          *string `json:"title,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCollectorConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCollectorConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowCollectorConnectionResponse", string(data)}, " ")
}
