package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDependencyVersionRequest Request Object
type ShowDependencyVersionRequest struct {

	// 依赖包的ID。
	DependId string `json:"depend_id"`

	// 依赖包版本号。
	Version string `json:"version"`
}

func (o ShowDependencyVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDependencyVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowDependencyVersionRequest", string(data)}, " ")
}
