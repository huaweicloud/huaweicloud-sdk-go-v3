package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询日志接入列表请求体
type GetAccessConfigListRequestBody struct {
	// 接入配置名称列表

	AccessConfigNameList []string `json:"access_config_name_list"`
	// 主机组名称列表

	HostGroupNameList []string `json:"host_group_name_list"`
	// 日志组名称列表

	LogGroupNameList []string `json:"log_group_name_list"`
	// 日志流名称列表

	LogStreamNameList []string `json:"log_stream_name_list"`

	AccessConfigTagList []AccessConfigTag `json:"access_config_tag_list"`
}

func (o GetAccessConfigListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAccessConfigListRequestBody struct{}"
	}

	return strings.Join([]string{"GetAccessConfigListRequestBody", string(data)}, " ")
}
