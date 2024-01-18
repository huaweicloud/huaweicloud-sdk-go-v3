package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataDetailResponse Response Object
type ShowDataDetailResponse struct {
	Entity *Entity `json:"entity,omitempty"`

	// 关联资产信息，结构体Map<String, Entity>
	ReferredEntities *interface{} `json:"referred_entities,omitempty"`
	HttpStatusCode   int          `json:"-"`
}

func (o ShowDataDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDataDetailResponse", string(data)}, " ")
}
