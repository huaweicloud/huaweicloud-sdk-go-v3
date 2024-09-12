package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchWriteKvResponse Response Object
type BatchWriteKvResponse struct {

	// 未处理的操作列表。
	UnprocessedOpers *[]TableOperIds `bson:"unprocessed_opers,omitempty"`
	HttpStatusCode   int             `bson:"-"`
}

func (o BatchWriteKvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchWriteKvResponse struct{}"
	}

	return strings.Join([]string{"BatchWriteKvResponse", string(data)}, " ")
}
