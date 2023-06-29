package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDictionaryRequest Request Object
type UpdateDictionaryRequest struct {

	// 字典ID
	DictId string `json:"dict_id"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *UpdateDictionary `json:"body,omitempty"`
}

func (o UpdateDictionaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDictionaryRequest struct{}"
	}

	return strings.Join([]string{"UpdateDictionaryRequest", string(data)}, " ")
}
