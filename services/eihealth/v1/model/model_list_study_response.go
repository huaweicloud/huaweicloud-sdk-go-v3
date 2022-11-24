package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStudyResponse struct {

	// study总数
	Count *int32 `json:"count,omitempty"`

	// study列表
	Studies        *[]StudyRsp `json:"studies,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListStudyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStudyResponse struct{}"
	}

	return strings.Join([]string{"ListStudyResponse", string(data)}, " ")
}
