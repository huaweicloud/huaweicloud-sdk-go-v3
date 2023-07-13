package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttrPair
type AttrPair struct {

	// 被推荐对象的属性名。
	PartyA *string `json:"party_a,omitempty"`

	// 被推荐对象的属性名。
	PartyB *string `json:"party_b,omitempty"`
}

func (o AttrPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttrPair struct{}"
	}

	return strings.Join([]string{"AttrPair", string(data)}, " ")
}
