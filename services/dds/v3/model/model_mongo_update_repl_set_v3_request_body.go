package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MongoUpdateReplSetV3RequestBody struct {

	// 连接地址复制集名称：实例高可用连接地址的唯一标识。该参数可以将读请求发送到副本集实例的所有节点。副本集中的所有主机必须具有相同的集名称。字符限制：大小写字母，数字，下划线组合，字母为首，长度限制在3-128
	Name string `json:"name"`
}

func (o MongoUpdateReplSetV3RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MongoUpdateReplSetV3RequestBody struct{}"
	}

	return strings.Join([]string{"MongoUpdateReplSetV3RequestBody", string(data)}, " ")
}
