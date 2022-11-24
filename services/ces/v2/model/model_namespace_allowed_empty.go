package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询服务的命名空间，各服务命名空间请参考[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)
type NamespaceAllowedEmpty struct {
}

func (o NamespaceAllowedEmpty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceAllowedEmpty struct{}"
	}

	return strings.Join([]string{"NamespaceAllowedEmpty", string(data)}, " ")
}
