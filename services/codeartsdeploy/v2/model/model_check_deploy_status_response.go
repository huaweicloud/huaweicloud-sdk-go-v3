package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDeployStatusResponse Response Object
type CheckDeployStatusResponse struct {
	Result *StateInfo `json:"result,omitempty"`

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckDeployStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDeployStatusResponse struct{}"
	}

	return strings.Join([]string{"CheckDeployStatusResponse", string(data)}, " ")
}
