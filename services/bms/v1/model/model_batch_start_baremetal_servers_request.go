package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchStartBaremetalServersRequest struct {
	Body *OsStartBody `json:"body,omitempty" xml:"body"`
}

func (o BatchStartBaremetalServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartBaremetalServersRequest struct{}"
	}

	return strings.Join([]string{"BatchStartBaremetalServersRequest", string(data)}, " ")
}
