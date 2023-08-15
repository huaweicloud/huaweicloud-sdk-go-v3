package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEntityInfoByGuidResponse Response Object
type ShowEntityInfoByGuidResponse struct {
	Entity *OpenEntityWithExtInfoEntity `json:"entity,omitempty"`

	// 引用实体 Map<String, OpenEntity>
	ReferredEntities *interface{} `json:"referred_entities,omitempty"`
	HttpStatusCode   int          `json:"-"`
}

func (o ShowEntityInfoByGuidResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEntityInfoByGuidResponse struct{}"
	}

	return strings.Join([]string{"ShowEntityInfoByGuidResponse", string(data)}, " ")
}
