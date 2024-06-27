package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIteratorRequest Request Object
type UpdateIteratorRequest struct {

	// 迭代uri
	IteratorId string `json:"iterator_id"`

	Body *IteratorVersionInfo `json:"body,omitempty"`
}

func (o UpdateIteratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIteratorRequest struct{}"
	}

	return strings.Join([]string{"UpdateIteratorRequest", string(data)}, " ")
}
