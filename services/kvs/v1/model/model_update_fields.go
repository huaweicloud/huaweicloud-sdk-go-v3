package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

type UpdateFields struct {

	// 文档不存在时，插入该完整文档。
	Upsert *bson.D `bson:"upsert,omitempty"`

	// 新增或覆盖更新1个或多个字段的值。 > 禁止修改sortkey的字段。
	Set *bson.D `bson:"set,omitempty"`

	// 对1个或多个字段做加法运算，并更新为运算后的值。
	Add *bson.D `bson:"add,omitempty"`

	// 删除1个或多个字段。 - 数组元素为待删除字段名。
	Rmv *[]string `bson:"rmv,omitempty"`

	// 插入元素到数组中 > 非数组返回失败。
	Insert *bson.D `bson:"insert,omitempty"`
}

func (o UpdateFields) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFields struct{}"
	}

	return strings.Join([]string{"UpdateFields", string(data)}, " ")
}
