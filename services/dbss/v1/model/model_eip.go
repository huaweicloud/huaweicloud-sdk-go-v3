package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Eip EIP信息
type Eip struct {
	Bandwidth *BandWidth `json:"bandwidth,omitempty"`

	// IP产品ID
	Ipproductid *string `json:"ipproductid,omitempty"`

	// EIP类型
	Iptype *string `json:"iptype,omitempty"`
}

func (o Eip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Eip struct{}"
	}

	return strings.Join([]string{"Eip", string(data)}, " ")
}
