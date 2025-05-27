package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchGetKvResponse Response Object
type BatchGetKvResponse struct {

	// 异常处理的操作，按照table分类组织。
	ExceptionOpers *[]ExceptionOpersOfTable `bson:"exception_opers,omitempty"`

	// 返回的kvdoc列表，按照table分类组织。
	ReturnedKvItemsOfAll *[]ReturnedKvItemsOfTable `bson:"returned_kv_items_of_all,omitempty"`
	HttpStatusCode       int                       `bson:"-"`
}

func (o BatchGetKvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchGetKvResponse struct{}"
	}

	return strings.Join([]string{"BatchGetKvResponse", string(data)}, " ")
}
