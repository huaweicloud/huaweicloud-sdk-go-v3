package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestorePtrRecordRequest struct {
	Region string `json:"region"`

	FloatingipId string `json:"floatingip_id"`

	Body *RestorePtrReq `json:"body,omitempty"`
}

func (o RestorePtrRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorePtrRecordRequest struct{}"
	}

	return strings.Join([]string{"RestorePtrRecordRequest", string(data)}, " ")
}
