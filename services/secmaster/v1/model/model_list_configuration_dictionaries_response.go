package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationDictionariesResponse Response Object
type ListConfigurationDictionariesResponse struct {

	// 正常字典列表
	SuccessList *[]Dictionary `json:"success_list,omitempty"`

	// 正常字典列表
	FailedList     *[]Dictionary `json:"failed_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListConfigurationDictionariesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationDictionariesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationDictionariesResponse", string(data)}, " ")
}
