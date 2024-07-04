package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckStarrocksParamsRequestBody struct {

	// 需要进行比较的源参数模板ID。通过ListStarrocksInstanceInfo接口获得。
	SourceConfigurationId string `json:"source_configuration_id"`
}

func (o CheckStarrocksParamsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckStarrocksParamsRequestBody struct{}"
	}

	return strings.Join([]string{"CheckStarrocksParamsRequestBody", string(data)}, " ")
}
