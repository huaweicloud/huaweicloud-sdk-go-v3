package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BankReceiptKvPair struct {

	// key-value对（键值对）中的key，例如\"币别：人民币\"中的\"币别\"。
	Key *string `json:"key,omitempty"`

	// key-value对（键值对）中的value，例如\"币别：人民币\"中的“人民币”
	Value *string `json:"value,omitempty"`
}

func (o BankReceiptKvPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BankReceiptKvPair struct{}"
	}

	return strings.Join([]string{"BankReceiptKvPair", string(data)}, " ")
}
