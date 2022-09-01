package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloseMysqlProxyRequestBody struct {

	// 数据库代理id列表。如果实例只开启了一个代理，可不传该字段；如果实例开启了多个代理，则必须指定要关闭的代理。
	ProxyIds *[]string `json:"proxy_ids,omitempty" xml:"proxy_ids"`
}

func (o CloseMysqlProxyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseMysqlProxyRequestBody struct{}"
	}

	return strings.Join([]string{"CloseMysqlProxyRequestBody", string(data)}, " ")
}
