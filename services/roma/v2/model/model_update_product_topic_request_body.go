package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateProductTopicRequestBody struct {
	// 主题权限 0-发布 1-订阅

	Permission int32 `json:"permission"`
	// 主题名称，Topic类格式必须以“/”进行分层，区分每个类目。其中第一个为用户自定义的版本号；第二个已经规定好，为${deviceId}通配设备ID；第三个为用户自定义的topic类名。Topic类组成即为：/${version}/${deviceId}/${customizePart}。简单来说，Topic类：/v1/${deviceId}/customizePart是具体Topic：/v1/deviceid1/customizePart1和/v1/deviceid2/customizePart2等的集合。

	Name string `json:"name"`
	// 版本号,输入2-50个字符。值以字母或数字开头和结尾。仅允许使用字母，数字，中划线和点号。

	Version string `json:"version"`
	// 描述，长度0-200

	Description *string `json:"description,omitempty"`
}

func (o UpdateProductTopicRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProductTopicRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateProductTopicRequestBody", string(data)}, " ")
}
