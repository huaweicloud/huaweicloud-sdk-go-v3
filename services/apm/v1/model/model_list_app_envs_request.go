package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAppEnvsRequest struct {

	// 组件id。
	AppId int64 `json:"app_id"`

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`
}

func (o ListAppEnvsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppEnvsRequest struct{}"
	}

	return strings.Join([]string{"ListAppEnvsRequest", string(data)}, " ")
}
