package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkTableIssuseListResponseBodyDomain 领域信息
type WorkTableIssuseListResponseBodyDomain struct {

	// 领域id
	Id *int32 `json:"id,omitempty"`

	// 领域名称
	Name *string `json:"name,omitempty"`
}

func (o WorkTableIssuseListResponseBodyDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkTableIssuseListResponseBodyDomain struct{}"
	}

	return strings.Join([]string{"WorkTableIssuseListResponseBodyDomain", string(data)}, " ")
}
