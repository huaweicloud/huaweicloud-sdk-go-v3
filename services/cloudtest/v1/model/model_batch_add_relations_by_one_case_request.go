package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddRelationsByOneCaseRequest Request Object
type BatchAddRelationsByOneCaseRequest struct {

	// 需求/缺陷id
	WorkitemId string `json:"workitem_id"`

	Body *AddRelationsInfo `json:"body,omitempty"`
}

func (o BatchAddRelationsByOneCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddRelationsByOneCaseRequest struct{}"
	}

	return strings.Join([]string{"BatchAddRelationsByOneCaseRequest", string(data)}, " ")
}
