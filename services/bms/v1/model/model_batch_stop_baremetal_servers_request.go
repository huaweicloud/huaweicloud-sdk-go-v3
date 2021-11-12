package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchStopBaremetalServersRequest struct {
	Body *OsStopBody `json:"body,omitempty"`
}

func (o BatchStopBaremetalServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopBaremetalServersRequest struct{}"
	}

	return strings.Join([]string{"BatchStopBaremetalServersRequest", string(data)}, " ")
}
