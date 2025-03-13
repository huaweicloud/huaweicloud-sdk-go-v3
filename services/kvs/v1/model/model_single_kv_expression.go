package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SingleKvExpression struct {

	// 取值：\"is_doc\", \"is_blob\", \"is_exist\", \"not_exist\"。
	Func string `bson:"func"`
}

func (o SingleKvExpression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SingleKvExpression struct{}"
	}

	return strings.Join([]string{"SingleKvExpression", string(data)}, " ")
}
