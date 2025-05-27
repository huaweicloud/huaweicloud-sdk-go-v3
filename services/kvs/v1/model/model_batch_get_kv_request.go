package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchGetKvRequest Request Object
type BatchGetKvRequest struct {

	// 仓名
	StoreName *string `bson:"store_name,omitempty"`

	Body *BatchGetKvRequestBody `bson:"body,omitempty"`
}

func (o BatchGetKvRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchGetKvRequest struct{}"
	}

	return strings.Join([]string{"BatchGetKvRequest", string(data)}, " ")
}
