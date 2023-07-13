package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutOfficialEventsRequest Request Object
type PutOfficialEventsRequest struct {

	// 事件源名称
	SourceName string `json:"source_name"`

	Body *PutEventsReq `json:"body,omitempty"`
}

func (o PutOfficialEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutOfficialEventsRequest struct{}"
	}

	return strings.Join([]string{"PutOfficialEventsRequest", string(data)}, " ")
}
