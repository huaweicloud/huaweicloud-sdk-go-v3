package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceAgencyRequest Request Object
type ListServiceAgencyRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListServiceAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceAgencyRequest struct{}"
	}

	return strings.Join([]string{"ListServiceAgencyRequest", string(data)}, " ")
}
