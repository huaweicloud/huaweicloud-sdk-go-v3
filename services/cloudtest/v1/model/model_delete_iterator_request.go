package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIteratorRequest Request Object
type DeleteIteratorRequest struct {

	// 迭代URI
	IteratorUri string `json:"iterator_uri"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 是否异步
	IsAsync *bool `json:"is_async,omitempty"`
}

func (o DeleteIteratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIteratorRequest struct{}"
	}

	return strings.Join([]string{"DeleteIteratorRequest", string(data)}, " ")
}
