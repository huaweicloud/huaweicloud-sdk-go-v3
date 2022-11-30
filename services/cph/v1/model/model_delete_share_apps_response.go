package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteShareAppsResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务列表。
	Jobs           *[]ServerJob `json:"jobs,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteShareAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareAppsResponse struct{}"
	}

	return strings.Join([]string{"DeleteShareAppsResponse", string(data)}, " ")
}
