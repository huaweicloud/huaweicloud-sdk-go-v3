package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnterpriseProjectResponse Response Object
type ShowEnterpriseProjectResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 企业项目列表。
	EnterpriseProjectList *[]EnterpriseProject `json:"enterprise_project_list,omitempty"`
	HttpStatusCode        int                  `json:"-"`
}

func (o ShowEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseProjectResponse", string(data)}, " ")
}
