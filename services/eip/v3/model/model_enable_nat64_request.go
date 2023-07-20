package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableNat64Request Request Object
type EnableNat64Request struct {

	// 弹性公网ID
	PublicipId string `json:"publicip_id"`
}

func (o EnableNat64Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableNat64Request struct{}"
	}

	return strings.Join([]string{"EnableNat64Request", string(data)}, " ")
}
