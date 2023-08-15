package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertificateAuthorityObsBucketResponse Response Object
type ListCertificateAuthorityObsBucketResponse struct {

	// OBS桶总数。
	Total *int32 `json:"total,omitempty"`

	// OBS桶列表，详情请参见**ObsBuckets**字段数据结构说明。
	ObsBuckets     *[]ObsBuckets `json:"obs_buckets,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListCertificateAuthorityObsBucketResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificateAuthorityObsBucketResponse struct{}"
	}

	return strings.Join([]string{"ListCertificateAuthorityObsBucketResponse", string(data)}, " ")
}
