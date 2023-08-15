package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainRequest Request Object
type ShowDomainRequest struct {

	// 源项目ID
	SourceProjectId string `json:"source_project_id"`
}

func (o ShowDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainRequest", string(data)}, " ")
}
