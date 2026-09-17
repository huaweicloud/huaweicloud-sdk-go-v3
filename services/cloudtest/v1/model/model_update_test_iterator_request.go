package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTestIteratorRequest Request Object
type UpdateTestIteratorRequest struct {

	// 迭代URI
	IteratorUri string `json:"iterator_uri"`

	Body *IteratorVersionInfo `json:"body,omitempty"`
}

func (o UpdateTestIteratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestIteratorRequest struct{}"
	}

	return strings.Join([]string{"UpdateTestIteratorRequest", string(data)}, " ")
}
