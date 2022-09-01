package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DIS数据源配置内容
type DisContentRsp struct {

	// 通道名称
	StreamName *string `json:"streamName,omitempty" xml:"streamName"`

	// 租户的AK
	Ak *string `json:"ak,omitempty" xml:"ak"`

	// 租户的SK
	Sk *string `json:"sk,omitempty" xml:"sk"`

	// 项目id
	ProjectId *string `json:"projectId,omitempty" xml:"projectId"`
}

func (o DisContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisContentRsp struct{}"
	}

	return strings.Join([]string{"DisContentRsp", string(data)}, " ")
}
