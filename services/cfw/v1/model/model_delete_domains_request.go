package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainsRequest Request Object
type DeleteDomainsRequest struct {

	// 域名组id
	SetId string `json:"set_id"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *DeleteDomainDto `json:"body,omitempty"`
}

func (o DeleteDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainsRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainsRequest", string(data)}, " ")
}
