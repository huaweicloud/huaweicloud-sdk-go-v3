package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandleCocIncidentRequest Request Object
type HandleCocIncidentRequest struct {
	Body *HandleExternalIncidentRequest `json:"body,omitempty"`
}

func (o HandleCocIncidentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleCocIncidentRequest struct{}"
	}

	return strings.Join([]string{"HandleCocIncidentRequest", string(data)}, " ")
}
