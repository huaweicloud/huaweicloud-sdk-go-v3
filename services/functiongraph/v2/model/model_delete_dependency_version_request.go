package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDependencyVersionRequest Request Object
type DeleteDependencyVersionRequest struct {

	// 依赖包的ID。
	DependId string `json:"depend_id"`

	// 依赖包版本号。
	Version string `json:"version"`
}

func (o DeleteDependencyVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDependencyVersionRequest struct{}"
	}

	return strings.Join([]string{"DeleteDependencyVersionRequest", string(data)}, " ")
}
