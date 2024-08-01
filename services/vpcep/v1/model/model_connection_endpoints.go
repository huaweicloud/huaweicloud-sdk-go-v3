package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionEndpoints 终端节点列表
type ConnectionEndpoints struct {

	// 终端节点的ID，唯一标识。
	Id *string `json:"id,omitempty"`

	// 终端节点的报文标识。
	MarkerId *int32 `json:"marker_id,omitempty"`

	// 终端节点的创建时间。 采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 终端节点的更新时间。 采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 用户的Domain ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 错误信息。  当终端节点服务状态异常，即“status”的值为“failed”时，会返回该字段。
	Error *[]QueryError `json:"error,omitempty"`

	// 终端节点的连接状态。  - pendingAcceptance：待接受  - creating：创建中  - accepted：已接受  - rejected：已拒绝  - failed：失败  - deleting：删除中
	Status *string `json:"status,omitempty"`

	// 终端节点连接描述。
	Description *string `json:"description,omitempty"`
}

func (o ConnectionEndpoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionEndpoints struct{}"
	}

	return strings.Join([]string{"ConnectionEndpoints", string(data)}, " ")
}
