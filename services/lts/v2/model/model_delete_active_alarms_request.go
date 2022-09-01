package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteActiveAlarmsRequest struct {

	// domainId
	DomainId string `json:"domain_id" xml:"domain_id"`

	Body *DeleteActiveAlarmsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DeleteActiveAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteActiveAlarmsRequest struct{}"
	}

	return strings.Join([]string{"DeleteActiveAlarmsRequest", string(data)}, " ")
}
