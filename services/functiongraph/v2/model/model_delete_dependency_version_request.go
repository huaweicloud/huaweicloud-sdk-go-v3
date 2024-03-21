package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDependencyVersionRequest Request Object
type DeleteDependencyVersionRequest struct {

	// 依赖包的ID。
	DependId string `json:"depend_id"`

	// 依赖包版本号。
	Version string `json:"version"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o DeleteDependencyVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDependencyVersionRequest struct{}"
	}

	return strings.Join([]string{"DeleteDependencyVersionRequest", string(data)}, " ")
}
