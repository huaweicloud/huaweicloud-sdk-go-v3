package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SuccessResources 带宽更新成功对象
type SuccessResources struct {

	// - 功能说明：更新成功的带宽id
	Id *string `json:"id,omitempty"`
}

func (o SuccessResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SuccessResources struct{}"
	}

	return strings.Join([]string{"SuccessResources", string(data)}, " ")
}
