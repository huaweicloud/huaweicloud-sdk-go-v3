package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskResultCommitResponse Response Object
type ShowTaskResultCommitResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowTaskResultCommitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskResultCommitResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskResultCommitResponse", string(data)}, " ")
}
