package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachShareFilesystemResponse Response Object
type DetachShareFilesystemResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务列表。
	Jobs           *[]AttachShareFilesystemResponseBody200Jobs `json:"jobs,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o DetachShareFilesystemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachShareFilesystemResponse struct{}"
	}

	return strings.Join([]string{"DetachShareFilesystemResponse", string(data)}, " ")
}
