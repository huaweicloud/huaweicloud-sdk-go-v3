package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateModulesResponse Response Object
type ListPrivateModulesResponse struct {

	// 私有模块的列表。默认以创建时间升序排序。
	Modules        *[]PrivateModuleSummary `json:"modules,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListPrivateModulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateModulesResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateModulesResponse", string(data)}, " ")
}
