package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRecyclePolicyResponse Response Object
type SetRecyclePolicyResponse struct {

	// 修改结果，“success”表示修改成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetRecyclePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRecyclePolicyResponse struct{}"
	}

	return strings.Join([]string{"SetRecyclePolicyResponse", string(data)}, " ")
}
