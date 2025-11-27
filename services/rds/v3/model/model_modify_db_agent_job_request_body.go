package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDbAgentJobRequestBody 修改数据库代理作业请求体。
type ModifyDbAgentJobRequestBody struct {

	// 配置文件id。
	ProfileId string `json:"profile_id"`
}

func (o ModifyDbAgentJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDbAgentJobRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyDbAgentJobRequestBody", string(data)}, " ")
}
