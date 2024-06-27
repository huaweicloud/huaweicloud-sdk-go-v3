package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIteratorDetailRequest Request Object
type ShowIteratorDetailRequest struct {

	// 迭代uri
	IteratorId string `json:"iterator_id"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`
}

func (o ShowIteratorDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIteratorDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowIteratorDetailRequest", string(data)}, " ")
}
