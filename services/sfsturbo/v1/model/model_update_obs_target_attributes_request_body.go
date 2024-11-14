package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateObsTargetAttributesRequestBody 更新后端存储属性请求体
type UpdateObsTargetAttributesRequestBody struct {
	Attributes *ObsTargetAttributes `json:"attributes"`
}

func (o UpdateObsTargetAttributesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateObsTargetAttributesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateObsTargetAttributesRequestBody", string(data)}, " ")
}
