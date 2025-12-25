package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCollectorChannelResponse Response Object
type ShowCollectorChannelResponse struct {

	// Iam用户ID
	CreateBy *string `json:"create_by,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	Error *ChannelErrorType `json:"error,omitempty"`

	// UUID
	GroupId *string `json:"group_id,omitempty"`

	// 相关描述信息
	Input *[]ShowModuleVo `json:"input,omitempty"`

	// 相关描述信息
	Nodes *[]NodeVo `json:"nodes,omitempty"`

	OperationStatus *ChannelOperationStatus `json:"operation_status,omitempty"`

	// 相关描述信息
	Output *[]ShowModuleVo `json:"output,omitempty"`

	// UUID
	ParserId *string `json:"parser_id,omitempty"`

	// 所属租户名称
	ParserName *string `json:"parser_name,omitempty"`

	// 相关标题
	Title          *string `json:"title,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCollectorChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCollectorChannelResponse struct{}"
	}

	return strings.Join([]string{"ShowCollectorChannelResponse", string(data)}, " ")
}
