package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MicroServiceInfoCceServiceBase CCE云容器引擎Service信息
type MicroServiceInfoCceServiceBase struct {

	// 云容器引擎集群编号
	ClusterId string `json:"cluster_id"`

	// 命名空间。1-63字符。只能包含小写字母、数字，以及 '-'，必须以字母开头，必须以字母数字结尾。
	Namespace string `json:"namespace"`

	// Service名称。支持汉字，英文，数字，点，中划线，下划线，且只能以英文和汉字开头，1-64字符。 > 中文字符必须为UTF-8或者unicode编码。
	ServiceName string `json:"service_name"`
}

func (o MicroServiceInfoCceServiceBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroServiceInfoCceServiceBase struct{}"
	}

	return strings.Join([]string{"MicroServiceInfoCceServiceBase", string(data)}, " ")
}
