package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExternalIncidentResponseData CreateExternalIncidentResponseData
type CreateExternalIncidentResponseData struct {

	// 事件单号
	IncidentNum string `json:"incident_num"`
}

func (o CreateExternalIncidentResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExternalIncidentResponseData struct{}"
	}

	return strings.Join([]string{"CreateExternalIncidentResponseData", string(data)}, " ")
}
