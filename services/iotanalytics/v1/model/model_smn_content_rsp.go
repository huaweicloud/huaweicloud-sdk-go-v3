package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmnContentRsp SMN数据源配置内容
type SmnContentRsp struct {

	// 项目id
	ProjectId *string `json:"projectId,omitempty"`

	// 租户的AK
	Ak *string `json:"ak,omitempty"`

	// 租户的SK
	Sk *string `json:"sk,omitempty"`
}

func (o SmnContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnContentRsp struct{}"
	}

	return strings.Join([]string{"SmnContentRsp", string(data)}, " ")
}
