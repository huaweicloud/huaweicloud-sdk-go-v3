package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConfigurationDictionariesResponse Response Object
type UpdateConfigurationDictionariesResponse struct {

	// 正常字典列表
	SuccessList *[]Dictionary `json:"success_list,omitempty"`

	// 正常字典列表
	FailedList     *[]Dictionary `json:"failed_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateConfigurationDictionariesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationDictionariesResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationDictionariesResponse", string(data)}, " ")
}
