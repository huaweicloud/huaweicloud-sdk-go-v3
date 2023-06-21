package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CCE云容器引擎Service信息，service_type为CCE_SERVICE时必填
type MicroServiceInfoCceServiceCreate struct {

	// 云容器引擎集群编号
	ClusterId string `json:"cluster_id"`

	// 命名空间。1-63字符。只能包含小写字母、数字，以及 '-'，必须以字母开头，必须以字母数字结尾。
	Namespace string `json:"namespace"`

	// Service名称。支持汉字，英文，数字，点，中划线，下划线，且只能以英文和汉字开头，1-64字符。 > 中文字符必须为UTF-8或者unicode编码。
	ServiceName string `json:"service_name"`

	// Service的监听端口号。如果是多端口Service，用户需填写一个端口。
	Port *int32 `json:"port,omitempty"`
}

func (o MicroServiceInfoCceServiceCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroServiceInfoCceServiceCreate struct{}"
	}

	return strings.Join([]string{"MicroServiceInfoCceServiceCreate", string(data)}, " ")
}
