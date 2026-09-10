package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachShareFilesystemResponse Response Object
type AttachShareFilesystemResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务列表。
	Jobs           *[]AttachShareFilesystemResponseBody200Jobs `json:"jobs,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o AttachShareFilesystemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachShareFilesystemResponse struct{}"
	}

	return strings.Join([]string{"AttachShareFilesystemResponse", string(data)}, " ")
}
