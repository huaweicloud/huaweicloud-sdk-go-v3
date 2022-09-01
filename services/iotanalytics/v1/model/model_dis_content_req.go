package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DIS数据源配置内容
type DisContentReq struct {

	// 项目id
	ProjectId string `json:"project_id" xml:"project_id"`

	// 通道名称
	StreamName string `json:"stream_name" xml:"stream_name"`

	// 租户的AK
	Ak string `json:"ak" xml:"ak"`

	// 租户的SK
	Sk string `json:"sk" xml:"sk"`
}

func (o DisContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisContentReq struct{}"
	}

	return strings.Join([]string{"DisContentReq", string(data)}, " ")
}
