package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisContentRsp DIS数据源配置内容
type DisContentRsp struct {

	// 通道名称
	StreamName *string `json:"streamName,omitempty"`

	// 租户的AK
	Ak *string `json:"ak,omitempty"`

	// 租户的SK
	Sk *string `json:"sk,omitempty"`

	// 项目id
	ProjectId *string `json:"projectId,omitempty"`
}

func (o DisContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisContentRsp struct{}"
	}

	return strings.Join([]string{"DisContentRsp", string(data)}, " ")
}
