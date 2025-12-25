package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectorModuleTemplateResponse Response Object
type ListCollectorModuleTemplateResponse struct {

	// 常用解析模板数组
	Common *[]ModuleTemplateVo `json:"common,omitempty"`

	// 列出解析模板数组
	List           *[]ModuleTemplateVo `json:"list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListCollectorModuleTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorModuleTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListCollectorModuleTemplateResponse", string(data)}, " ")
}
