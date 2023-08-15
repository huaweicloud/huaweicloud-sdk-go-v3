package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDuplicateApisForAppV2Request Request Object
type ListDuplicateApisForAppV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 应用id
	AppId string `json:"app_id"`
}

func (o ListDuplicateApisForAppV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDuplicateApisForAppV2Request struct{}"
	}

	return strings.Join([]string{"ListDuplicateApisForAppV2Request", string(data)}, " ")
}
