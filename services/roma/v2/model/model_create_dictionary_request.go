package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDictionaryRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *CreateDictionary `json:"body,omitempty" xml:"body"`
}

func (o CreateDictionaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDictionaryRequest struct{}"
	}

	return strings.Join([]string{"CreateDictionaryRequest", string(data)}, " ")
}
