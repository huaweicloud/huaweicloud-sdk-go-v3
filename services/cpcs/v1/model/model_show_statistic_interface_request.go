package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatisticInterfaceRequest Request Object
type ShowStatisticInterfaceRequest struct {

	// 集群id，默认为空，默认查询所有集群
	ClusterId *string `json:"cluster_id,omitempty"`

	// 应用id，默认为空，默认查询所有app
	AppId *string `json:"app_id,omitempty"`

	// 查询的初始时间戳，毫秒级时间戳，默认查询前三天
	From *int64 `json:"from,omitempty"`

	// 查询的终止时间戳，毫秒级时间戳，默认查询到当前时间
	To *int64 `json:"to,omitempty"`

	// 统计周期，默认为1，五分钟为一个周期
	Period *int32 `json:"period,omitempty"`

	// 统计值，默认为min:最小值
	Filter *string `json:"filter,omitempty"`

	// 算法，有：“sm2”,\"rsa\"
	Algorithm *string `json:"algorithm,omitempty"`

	// 算法类型，0：国密算法，1：国际算法
	AlgorithmType *string `json:"algorithm_type,omitempty"`

	// 证书类型，0：根证书，1：业务证书
	CertificateType *string `json:"certificate_type,omitempty"`

	// 密码服务类型
	ServerType *string `json:"server_type,omitempty"`
}

func (o ShowStatisticInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticInterfaceRequest struct{}"
	}

	return strings.Join([]string{"ShowStatisticInterfaceRequest", string(data)}, " ")
}
