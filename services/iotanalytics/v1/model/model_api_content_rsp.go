package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApiContentRsp API数据源配置内容
type ApiContentRsp struct {

	// 数据源名称
	Name *string `json:"name,omitempty"`

	// 数据上报url
	Url *string `json:"url,omitempty"`
}

func (o ApiContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiContentRsp struct{}"
	}

	return strings.Join([]string{"ApiContentRsp", string(data)}, " ")
}
