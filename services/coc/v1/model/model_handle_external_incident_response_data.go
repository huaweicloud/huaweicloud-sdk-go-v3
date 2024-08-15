package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandleExternalIncidentResponseData HandleExternalIncidentResponseData
type HandleExternalIncidentResponseData struct {
}

func (o HandleExternalIncidentResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleExternalIncidentResponseData struct{}"
	}

	return strings.Join([]string{"HandleExternalIncidentResponseData", string(data)}, " ")
}
