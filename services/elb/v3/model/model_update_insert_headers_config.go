package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInsertHeadersConfig 参数解释：要添加的请求头参数列表。
type UpdateInsertHeadersConfig struct {

	// 参数解释：要添加请求头参数列表。
	Configs []UpdateInsertHeaderConfig `json:"configs"`
}

func (o UpdateInsertHeadersConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInsertHeadersConfig struct{}"
	}

	return strings.Join([]string{"UpdateInsertHeadersConfig", string(data)}, " ")
}
