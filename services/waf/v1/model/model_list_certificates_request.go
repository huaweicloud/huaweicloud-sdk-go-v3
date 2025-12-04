package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCertificatesRequest Request Object
type ListCertificatesRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 分页查询时，返回第几页数据。默认值为1，表示返回第1页数据。
	Page *int32 `json:"page,omitempty"`

	// 分页查询时，每页包含多少条结果。范围1-100，默认值为10，表示每页包含10条结果。
	Pagesize *int32 `json:"pagesize,omitempty"`

	// 证书名称
	Name *string `json:"name,omitempty"`

	// 是否获取证书关联的域名，默认为false   -true:获取已关联域名的证书   -false:获取未关联域名的证书
	Host *bool `json:"host,omitempty"`

	// **参数解释：** 证书过期状态 **约束限制：** 不涉及 **取值范围：**  - 0:未过期  - 1:已过期  - 2:即将过期（证书将在一个月内过期）  **默认取值：** 不涉及
	ExpStatus *ListCertificatesRequestExpStatus `json:"exp_status,omitempty"`

	// 查询结果的证书来源服务是否包括SCM服务，值为true或者false。
	QueryScm *bool `json:"query_scm,omitempty"`
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

func (c ListCertificatesRequestExpStatus) Value() int32 {
	return c.value
}

func (c ListCertificatesRequestExpStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCertificatesRequestExpStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
