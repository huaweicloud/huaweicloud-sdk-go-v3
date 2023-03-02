package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckPutEventsReq struct {
	Sources *[]CheckPutEventsReqSources `json:"sources,omitempty"`
}

func (o CheckPutEventsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPutEventsReq struct{}"
	}

	return strings.Join([]string{"CheckPutEventsReq", string(data)}, " ")
}
