package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListModulesResponse Response Object
type BatchListModulesResponse struct {

	// 总记录数
	Count *int64 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`

	// 每页记录数
	Modules        *[]EdgeModuleDto `json:"modules,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchListModulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListModulesResponse struct{}"
	}

	return strings.Join([]string{"BatchListModulesResponse", string(data)}, " ")
}
