package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchReinstallServerResponse Response Object
type BatchReinstallServerResponse struct {

	// 服务器任务信息。
	Items          *[]ServerJobInfo `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchReinstallServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchReinstallServerResponse struct{}"
	}

	return strings.Join([]string{"BatchReinstallServerResponse", string(data)}, " ")
}
