package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInsertHeadersConfig 参数解释：要添加的请求头参数列表。
type CreateInsertHeadersConfig struct {

	// 参数解释：要添加请求头参数列表。
	Configs []CreateInsertHeaderConfig `json:"configs"`
}

func (o CreateInsertHeadersConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInsertHeadersConfig struct{}"
	}

	return strings.Join([]string{"CreateInsertHeadersConfig", string(data)}, " ")
}
