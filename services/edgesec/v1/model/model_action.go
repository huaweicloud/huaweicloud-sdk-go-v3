package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Action 用来保存网站反爬虫中特征反爬虫的防护动作信息
type Action struct {

	// 特征反爬虫中防护动作信息   - log：仅记录   - block：拦截
	Category *string `json:"category,omitempty"`
}

func (o Action) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Action struct{}"
	}

	return strings.Join([]string{"Action", string(data)}, " ")
}
