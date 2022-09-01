package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 策略修改body
type PolicyUpdateReq struct {
	Policy *PolicyUpdate `json:"policy" xml:"policy"`
}

func (o PolicyUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyUpdateReq struct{}"
	}

	return strings.Join([]string{"PolicyUpdateReq", string(data)}, " ")
}
