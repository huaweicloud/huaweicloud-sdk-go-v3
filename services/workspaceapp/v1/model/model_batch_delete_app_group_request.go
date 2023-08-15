package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppGroupRequest Request Object
type BatchDeleteAppGroupRequest struct {
	Body *DeleteAppGroupReq `json:"body,omitempty"`
}

func (o BatchDeleteAppGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppGroupRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppGroupRequest", string(data)}, " ")
}
