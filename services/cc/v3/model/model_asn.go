package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Asn 网络实例BGP协议的AS号。
type Asn struct {

	// 网络实例BGP协议的AS号。
	Asn int64 `json:"asn"`
}

func (o Asn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Asn struct{}"
	}

	return strings.Join([]string{"Asn", string(data)}, " ")
}
