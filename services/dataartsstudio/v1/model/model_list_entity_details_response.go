package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEntityDetailsResponse Response Object
type ListEntityDetailsResponse struct {

	// 资产信息列表。
	Entities *[]Entity `json:"entities,omitempty"`

	// 关联资产信息，结构体Map<String, Entity>。key：资产guid，value：资产信息。
	ReferredEntities *interface{} `json:"referred_entities,omitempty"`
	HttpStatusCode   int          `json:"-"`
}

func (o ListEntityDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEntityDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListEntityDetailsResponse", string(data)}, " ")
}
