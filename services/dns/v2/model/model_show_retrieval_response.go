package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRetrievalResponse struct {

	// 找回请求ID。
	Id *string `json:"id,omitempty"`

	// 域名名称。
	ZoneName *string `json:"zone_name,omitempty"`

	Record *CreatePublicZoneFindRespRecord `json:"record,omitempty"`

	// 状态，(PENDING,VERIFIED,CREATED,EXPIRED)
	Status *string `json:"status,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRetrievalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRetrievalResponse struct{}"
	}

	return strings.Join([]string{"ShowRetrievalResponse", string(data)}, " ")
}
