package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainVisibleServicesRequest Request Object
type ListDomainVisibleServicesRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ListDomainVisibleServicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainVisibleServicesRequest struct{}"
	}

	return strings.Join([]string{"ListDomainVisibleServicesRequest", string(data)}, " ")
}
