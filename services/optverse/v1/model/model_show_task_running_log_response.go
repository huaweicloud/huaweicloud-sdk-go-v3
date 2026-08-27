package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskRunningLogResponse Response Object
type ShowTaskRunningLogResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowTaskRunningLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskRunningLogResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskRunningLogResponse", string(data)}, " ")
}
