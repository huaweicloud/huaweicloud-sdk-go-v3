package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIncidentRequestBody 创建事件请求body体
type CreateIncidentRequestBody struct {
	DataObject *Incident `json:"data_object,omitempty"`
}

func (o CreateIncidentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIncidentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIncidentRequestBody", string(data)}, " ")
}
