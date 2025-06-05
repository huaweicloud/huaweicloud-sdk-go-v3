package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProductdataOfferingInfosRequest Request Object
type ShowProductdataOfferingInfosRequest struct {

	// Region ID
	Region *string `json:"region,omitempty"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 站点信息：   - HWC_CN ：中国站   - HWC_HK ：国际站   - HWC_EU : 欧洲站
	SiteCode *string `json:"site_code,omitempty"`
}

func (o ShowProductdataOfferingInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProductdataOfferingInfosRequest struct{}"
	}

	return strings.Join([]string{"ShowProductdataOfferingInfosRequest", string(data)}, " ")
}
