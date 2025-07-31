package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectConfigsRequestBody 查询配置信息请求体
type ListProjectConfigsRequestBody struct {

	// 配置名称列表
	ConfigNameList []string `json:"config_name_list"`
}

func (o ListProjectConfigsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectConfigsRequestBody struct{}"
	}

	return strings.Join([]string{"ListProjectConfigsRequestBody", string(data)}, " ")
}
