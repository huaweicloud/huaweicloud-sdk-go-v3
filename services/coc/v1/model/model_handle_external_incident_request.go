package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandleExternalIncidentRequest HandleExternalIncidentRequest
type HandleExternalIncidentRequest struct {

	// 事件单号
	IncidentNum string `json:"incident_num"`

	// 操作人ID
	Operator string `json:"operator"`

	// 操作类型
	OperateKey string `json:"operate_key"`

	// 入参
	Parameter map[string]interface{} `json:"parameter,omitempty"`
}

func (o HandleExternalIncidentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleExternalIncidentRequest struct{}"
	}

	return strings.Join([]string{"HandleExternalIncidentRequest", string(data)}, " ")
}
