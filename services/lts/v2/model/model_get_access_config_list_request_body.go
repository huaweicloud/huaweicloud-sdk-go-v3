package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询日志接入列表请求体
type GetAccessConfigListRequestBody struct {
	// 接入配置名称列表

	AccessConfigNameList *[]string `json:"access_config_name_list,omitempty"`
	// 主机组名称列表

	HostGroupNameList *[]string `json:"host_group_name_list,omitempty"`
	// 日志组名称列表

	LogGroupNameList *[]string `json:"log_group_name_list,omitempty"`
	// 日志流名称列表

	LogStreamNameList *[]string `json:"log_stream_name_list,omitempty"`

	AccessConfigTagList *[]AccessConfigTag `json:"access_config_tag_list,omitempty"`
}

func (o GetAccessConfigListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAccessConfigListRequestBody struct{}"
	}

	return strings.Join([]string{"GetAccessConfigListRequestBody", string(data)}, " ")
}
