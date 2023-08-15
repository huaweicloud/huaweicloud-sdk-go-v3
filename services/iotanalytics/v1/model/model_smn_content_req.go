package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmnContentReq SMN数据源配置内容
type SmnContentReq struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 租户的AK
	Ak string `json:"ak"`

	// 租户的SK
	Sk string `json:"sk"`
}

func (o SmnContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnContentReq struct{}"
	}

	return strings.Join([]string{"SmnContentReq", string(data)}, " ")
}
