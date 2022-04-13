package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDependencyRequest struct {
	// 依赖包的ID。

	DependId string `json:"depend_id"`
}

func (o ShowDependencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDependencyRequest struct{}"
	}

	return strings.Join([]string{"ShowDependencyRequest", string(data)}, " ")
}
