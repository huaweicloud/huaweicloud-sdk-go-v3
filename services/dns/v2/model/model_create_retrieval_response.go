package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRetrievalResponse Response Object
type CreateRetrievalResponse struct {

	// 找回请求ID。
	Id *string `json:"id,omitempty"`

	// 域名。
	ZoneName *string `json:"zone_name,omitempty"`

	Record *RecordInfo `json:"record,omitempty"`

	// 找回状态。  取值范围：  PENDING：处理中 VERIFIED：已验证 CREATED：已找回 EXPIRED：已失效
	Status *string `json:"status,omitempty"`

	// 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRetrievalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetrievalResponse struct{}"
	}

	return strings.Join([]string{"CreateRetrievalResponse", string(data)}, " ")
}
