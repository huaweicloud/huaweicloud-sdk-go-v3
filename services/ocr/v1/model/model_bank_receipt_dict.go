package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BankReceiptDict struct {

	// 键值对数量
	KvPairCount *int32 `json:"kv_pair_count,omitempty"`

	// 银行回单的区域位置信息，列表形式，包含文字区域四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	BankReceiptLocation *[][]int32 `json:"bank_receipt_location,omitempty"`

	// 键值对识别结果列表。
	KvPairList *[]BankReceiptKvPair `json:"kv_pair_list,omitempty"`
}

func (o BankReceiptDict) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BankReceiptDict struct{}"
	}

	return strings.Join([]string{"BankReceiptDict", string(data)}, " ")
}
