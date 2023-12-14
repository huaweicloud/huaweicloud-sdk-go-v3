package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBaremetalServerRequest Request Object
type DeleteBaremetalServerRequest struct {
	Body *DeleteBaremetalBody `json:"body,omitempty"`
}

func (o DeleteBaremetalServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBaremetalServerRequest struct{}"
	}

	return strings.Join([]string{"DeleteBaremetalServerRequest", string(data)}, " ")
}
