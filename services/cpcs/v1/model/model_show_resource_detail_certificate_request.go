package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceDetailCertificateRequest Request Object
type ShowResourceDetailCertificateRequest struct {

	// 集群id，默认空字符串，默认查询所有集群
	ClusterId *string `json:"cluster_id,omitempty"`

	// 应用id，默认空字符串，默认查询所有的应用
	AppId *string `json:"app_id,omitempty"`

	// 密码服务类型，默认空字符串，默认查询所有密码服务类型
	ServiceType *string `json:"service_type,omitempty"`

	// 算法类型，默认空字符串，0：国密，1：国际
	AlgorithmType *string `json:"algorithm_type,omitempty"`

	// 证书类型，默认空字符串，0：根证书，1：业务证书
	CertificateType *string `json:"certificate_type,omitempty"`

	// 页面大小，不超过1500
	PageSize *int32 `json:"page_size,omitempty"`

	// 页数，默认1
	PageNum *int32 `json:"page_num,omitempty"`

	// 查询起始时间戳，毫秒级时间戳，默认为0，默认从三天前查询
	From *int64 `json:"from,omitempty"`

	// 查询终止时间戳，毫秒级时间戳，默认为0，默认查询到当前时间
	To *int64 `json:"to,omitempty"`
}

func (o ShowResourceDetailCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceDetailCertificateRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceDetailCertificateRequest", string(data)}, " ")
}
