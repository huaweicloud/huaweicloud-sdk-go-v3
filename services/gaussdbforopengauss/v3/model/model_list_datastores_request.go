package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatastoresRequest Request Object
type ListDatastoresRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListDatastoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoresRequest struct{}"
	}

	return strings.Join([]string{"ListDatastoresRequest", string(data)}, " ")
}
