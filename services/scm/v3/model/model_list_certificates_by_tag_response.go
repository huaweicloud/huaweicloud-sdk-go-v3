package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertificatesByTagResponse Response Object
type ListCertificatesByTagResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 资源实例列表，详情请参见TagResource字段数据结构说明。
	Resources      *[]TagResource `json:"resources,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListCertificatesByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesByTagResponse struct{}"
	}

	return strings.Join([]string{"ListCertificatesByTagResponse", string(data)}, " ")
}
