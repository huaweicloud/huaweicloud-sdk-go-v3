package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateDbDataRequest Request Object
type ValidateDbDataRequest struct {

	// cluster_id
	ClusterId string `json:"cluster_id"`

	// database
	Database string `json:"database"`

	Body *ValidateDbDataReq `json:"body,omitempty"`
}

func (o ValidateDbDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateDbDataRequest struct{}"
	}

	return strings.Join([]string{"ValidateDbDataRequest", string(data)}, " ")
}
