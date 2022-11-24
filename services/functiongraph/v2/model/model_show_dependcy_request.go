package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDependcyRequest struct {

	// 依赖包的ID。
	DependId string `json:"depend_id"`
}

func (o ShowDependcyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDependcyRequest struct{}"
	}

	return strings.Join([]string{"ShowDependcyRequest", string(data)}, " ")
}
