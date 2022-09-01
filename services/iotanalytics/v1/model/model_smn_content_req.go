package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SMN数据源配置内容
type SmnContentReq struct {

	// 项目id
	ProjectId string `json:"project_id" xml:"project_id"`

	// 租户的AK
	Ak string `json:"ak" xml:"ak"`

	// 租户的SK
	Sk string `json:"sk" xml:"sk"`
}

func (o SmnContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnContentReq struct{}"
	}

	return strings.Join([]string{"SmnContentReq", string(data)}, " ")
}
