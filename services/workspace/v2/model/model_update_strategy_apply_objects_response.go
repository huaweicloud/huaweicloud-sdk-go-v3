package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStrategyApplyObjectsResponse Response Object
type UpdateStrategyApplyObjectsResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateStrategyApplyObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStrategyApplyObjectsResponse struct{}"
	}

	return strings.Join([]string{"UpdateStrategyApplyObjectsResponse", string(data)}, " ")
}
