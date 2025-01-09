package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeServerImageResponse Response Object
type BatchChangeServerImageResponse struct {

	// 服务器任务信息。
	Items          *[]ServerJobInfo `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchChangeServerImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeServerImageResponse struct{}"
	}

	return strings.Join([]string{"BatchChangeServerImageResponse", string(data)}, " ")
}
