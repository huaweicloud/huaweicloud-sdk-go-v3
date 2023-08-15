package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeploymentTag struct {

	// 部署标签的key值，长度取值范围为1~36，由英文字母，数字，中划线和下划线组成，长度1到36个字符
	Key string `json:"key"`

	// 部署标签的value值，长度取值范围为1~43，由英文字母，数字，下划线，点号和中划线组成，长度0到43个字符
	Value string `json:"value"`
}

func (o DeploymentTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentTag struct{}"
	}

	return strings.Join([]string{"DeploymentTag", string(data)}, " ")
}
