package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListCertificatesRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 分页查询时，返回第几页数据。范围0-100000，默认值为1，表示返回第1页数据。

	Page *int32 `json:"page,omitempty"`
	// 分页查询时，每页包含多少条结果。范围1-100，默认值为10，表示每页包含10条结果。

	Pagesize *int32 `json:"pagesize,omitempty"`
	// 证书名称

	Name *string `json:"name,omitempty"`
	// 是否获取证书关联的域名，默认为false   -true:获取已关联域名的证书   -false:获取未关联域名的证书

	Host *bool `json:"host,omitempty"`
	// 证书过期状态，0-未过期，1-已过期，2-即将过期（证书将在一个月内过期）

	ExpStatus *ListCertificatesRequestExpStatus `json:"exp_status,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}

type ListCertificatesRequestExpStatus struct {
	value int32
}

type ListCertificatesRequestExpStatusEnum struct {
	E_0 ListCertificatesRequestExpStatus
	E_1 ListCertificatesRequestExpStatus
	E_2 ListCertificatesRequestExpStatus
}

func GetListCertificatesRequestExpStatusEnum() ListCertificatesRequestExpStatusEnum {
	return ListCertificatesRequestExpStatusEnum{
		E_0: ListCertificatesRequestExpStatus{
			value: 0,
		}, E_1: ListCertificatesRequestExpStatus{
			value: 1,
		}, E_2: ListCertificatesRequestExpStatus{
			value: 2,
		},
	}
}

func (c ListCertificatesRequestExpStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCertificatesRequestExpStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
