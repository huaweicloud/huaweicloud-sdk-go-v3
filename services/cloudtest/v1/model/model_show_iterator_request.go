package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIteratorRequest Request Object
type ShowIteratorRequest struct {

	// 迭代uri
	IteratorUri string `json:"iterator_uri"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`
}

func (o ShowIteratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIteratorRequest struct{}"
	}

	return strings.Join([]string{"ShowIteratorRequest", string(data)}, " ")
}
