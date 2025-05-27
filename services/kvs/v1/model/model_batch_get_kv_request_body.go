package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchGetKvRequestBody struct {

	// 按照table分类组织的get操作。
	BatchGetKvOpers []BatchGetKvOfTable `bson:"batch_get_kv_opers"`
}

func (o BatchGetKvRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchGetKvRequestBody struct{}"
	}

	return strings.Join([]string{"BatchGetKvRequestBody", string(data)}, " ")
}
