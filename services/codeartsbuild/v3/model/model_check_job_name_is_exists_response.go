package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckJobNameIsExistsResponse Response Object
type CheckJobNameIsExistsResponse struct {

	// 项目下任务名是否存在
	Result *bool `json:"result,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckJobNameIsExistsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckJobNameIsExistsResponse struct{}"
	}

	return strings.Join([]string{"CheckJobNameIsExistsResponse", string(data)}, " ")
}
