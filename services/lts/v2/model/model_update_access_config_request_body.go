package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改日志接入请求体
type UpdateAccessConfigRequestBody struct {
	// 日志接入ID

	AccessConfigId string `json:"access_config_id"`

	AccessConfigDetail *AccessConfigDeatil `json:"access_config_detail,omitempty"`

	HostGroupInfo *AccessConfigHostGroupIdList `json:"host_group_info,omitempty"`

	AccessConfigTag *[]AccessConfigTag `json:"access_config_tag,omitempty"`
}

func (o UpdateAccessConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAccessConfigRequestBody", string(data)}, " ")
}
