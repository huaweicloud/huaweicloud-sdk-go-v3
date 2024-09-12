package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KvOperIds struct {

	// 上传kv操作, \"oper_id\"数组。 - 数组元素：请求内的操作编码，未成功的操作返回该标识。
	PutKvIds *[]int32 `bson:"put_kv_ids,omitempty"`

	// 请求内的操作编码，未成功的操作返回该标识。 - 数组元素：请求内的操作编码，未成功的操作返回该标识。
	DeleteKvIds *[]int32 `bson:"delete_kv_ids,omitempty"`
}

func (o KvOperIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KvOperIds struct{}"
	}

	return strings.Join([]string{"KvOperIds", string(data)}, " ")
}
