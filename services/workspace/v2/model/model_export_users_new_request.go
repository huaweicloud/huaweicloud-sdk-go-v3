package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUsersNewRequest Request Object
type ExportUsersNewRequest struct {
	Body *ExportUsersNewReq `json:"body,omitempty"`
}

func (o ExportUsersNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUsersNewRequest struct{}"
	}

	return strings.Join([]string{"ExportUsersNewRequest", string(data)}, " ")
}
