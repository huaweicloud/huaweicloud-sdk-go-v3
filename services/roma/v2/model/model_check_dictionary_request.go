package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckDictionaryRequest struct {

	// 字典ID
	DictId string `json:"dict_id"`

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o CheckDictionaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDictionaryRequest struct{}"
	}

	return strings.Join([]string{"CheckDictionaryRequest", string(data)}, " ")
}
