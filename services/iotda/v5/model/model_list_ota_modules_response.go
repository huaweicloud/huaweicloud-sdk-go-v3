package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOtaModulesResponse Response Object
type ListOtaModulesResponse struct {

	// 模块列表
	Modules *[]OtaModuleInfo `json:"modules,omitempty"`

	Page           *PageInfo `json:"page,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListOtaModulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOtaModulesResponse struct{}"
	}

	return strings.Join([]string{"ListOtaModulesResponse", string(data)}, " ")
}
