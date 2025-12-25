package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigurationDictionariesResponse Response Object
type CreateConfigurationDictionariesResponse struct {

	// 正常字典列表
	SuccessList *[]Dictionary `json:"success_list,omitempty"`

	// 正常字典列表
	FailedList     *[]Dictionary `json:"failed_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateConfigurationDictionariesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationDictionariesResponse struct{}"
	}

	return strings.Join([]string{"CreateConfigurationDictionariesResponse", string(data)}, " ")
}
