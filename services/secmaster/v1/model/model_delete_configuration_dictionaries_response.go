package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConfigurationDictionariesResponse Response Object
type DeleteConfigurationDictionariesResponse struct {

	// 正常字典列表
	SuccessList *[]Dictionary `json:"success_list,omitempty"`

	// 正常字典列表
	FailedList     *[]Dictionary `json:"failed_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o DeleteConfigurationDictionariesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigurationDictionariesResponse struct{}"
	}

	return strings.Join([]string{"DeleteConfigurationDictionariesResponse", string(data)}, " ")
}
