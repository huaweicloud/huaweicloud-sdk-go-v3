package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResourceBindEnterpriseProjectResponse struct {
	// 资源列表

	Resources *[]Resources `json:"resources,omitempty"`
	// 查询失败的企业项目下的资源

	Errors *[]Errors `json:"errors,omitempty"`
	// 企业项目下的资源总数

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowResourceBindEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceBindEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceBindEnterpriseProjectResponse", string(data)}, " ")
}
