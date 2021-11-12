package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowOpLogResponse struct {
	OperationLog   *OperationLog `json:"operation_log,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowOpLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOpLogResponse struct{}"
	}

	return strings.Join([]string{"ShowOpLogResponse", string(data)}, " ")
}
