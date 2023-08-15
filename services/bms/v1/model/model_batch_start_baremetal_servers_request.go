package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartBaremetalServersRequest Request Object
type BatchStartBaremetalServersRequest struct {
	Body *OsStartBody `json:"body,omitempty"`
}

func (o BatchStartBaremetalServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartBaremetalServersRequest struct{}"
	}

	return strings.Join([]string{"BatchStartBaremetalServersRequest", string(data)}, " ")
}
