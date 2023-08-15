package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessConfigRequestBody 修改日志接入请求体
type UpdateAccessConfigRequestBody struct {

	// 日志接入ID
	AccessConfigId string `json:"access_config_id"`

	AccessConfigDetail *AccessConfigDeatilCreate `json:"access_config_detail,omitempty"`

	HostGroupInfo *AccessConfigHostGroupIdList `json:"host_group_info,omitempty"`

	// 标签信息。KEY不能重复,最多20个标签
	AccessConfigTag *[]AccessConfigTag `json:"access_config_tag,omitempty"`

	// 日志拆分
	LogSplit *bool `json:"log_split,omitempty"`

	// 二进制采集
	BinaryCollect *bool `json:"binary_collect,omitempty"`
}

func (o UpdateAccessConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAccessConfigRequestBody", string(data)}, " ")
}
