package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListCertificatesRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 页码

	Page *int32 `json:"page,omitempty"`
	// 每页条数

	Pagesize *int32 `json:"pagesize,omitempty"`
	// 证书名称

	Name *string `json:"name,omitempty"`
	// 是否获取证书关联的域名

	Host *bool `json:"host,omitempty"`
	// 证书过期状态，0-未过期，1-已过期，2-即将过期

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
