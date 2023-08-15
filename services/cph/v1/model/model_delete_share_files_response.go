package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteShareFilesResponse Response Object
type DeleteShareFilesResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务列表。
	Jobs           *[]ServerJob `json:"jobs,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteShareFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareFilesResponse struct{}"
	}

	return strings.Join([]string{"DeleteShareFilesResponse", string(data)}, " ")
}
