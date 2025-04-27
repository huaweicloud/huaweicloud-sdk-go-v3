package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPtrsResponse Response Object
type ListPtrsResponse struct {

	// 反向解析列表。
	Floatingips *[]FloatingIpsPtr `json:"floatingips,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPtrsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPtrsResponse struct{}"
	}

	return strings.Join([]string{"ListPtrsResponse", string(data)}, " ")
}
