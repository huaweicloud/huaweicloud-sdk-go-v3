package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateFeaturesRequestBody struct {

	// 需要修改的特性列表,参数值对象Map<String,String>。
	Params *interface{} `json:"params"`
}

func (o UpdateFeaturesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFeaturesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFeaturesRequestBody", string(data)}, " ")
}
