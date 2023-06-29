package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushShareFilesRequest Request Object
type PushShareFilesRequest struct {
	Body *PushShareFilesRequestBody `json:"body,omitempty"`
}

func (o PushShareFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushShareFilesRequest struct{}"
	}

	return strings.Join([]string{"PushShareFilesRequest", string(data)}, " ")
}
