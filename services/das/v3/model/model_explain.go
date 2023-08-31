package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Explain struct {

	// select子句的编号
	Id *int32 `json:"id,omitempty"`

	// select子句的类型
	SelectType *string `json:"select_type,omitempty"`

	// SQL优化器选择的表join顺序。
	Table *string `json:"table,omitempty"`

	// 查找表中行的访问类型(从好到坏依次为：null>system>const>eq_ref>ref>range>index>all)。
	Type *string `json:"type,omitempty"`

	// 有助于高效查找行的索引。
	PossibleKeys *string `json:"possible_keys,omitempty"`

	// 出于最小化查询成本考虑，SQL优化器实际使用的索引
	Key *string `json:"key,omitempty"`

	// key列所示索引的长度（字节）
	KeyLen *string `json:"key_len,omitempty"`

	// 在使用key列所示索引查找数据时用到的列或常量
	Ref *string `json:"ref,omitempty"`

	// key列所示索引的长度（字节）
	Rows *int64 `json:"rows,omitempty"`

	// sql解析的额外信息：当出现using index时，说明SQL使用覆盖索引，性能较好；而当出现 using filesort、using temporary、using where时，说明查询需要优化。
	Filtered *float64 `json:"filtered,omitempty"`

	// sql解析的额外信息：当出现using index时，说明SQL使用覆盖索引，性能较好；而当出现 using filesort、using temporary、using where时，说明查询需要优化。
	Extra *string `json:"extra,omitempty"`
}

func (o Explain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Explain struct{}"
	}

	return strings.Join([]string{"Explain", string(data)}, " ")
}
