package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckIsDuplicateAppNameResponse Response Object
type CheckIsDuplicateAppNameResponse struct {

	// 项目下是否存在同名应用
	Result *bool `json:"result,omitempty"`

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckIsDuplicateAppNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckIsDuplicateAppNameResponse struct{}"
	}

	return strings.Join([]string{"CheckIsDuplicateAppNameResponse", string(data)}, " ")
}
