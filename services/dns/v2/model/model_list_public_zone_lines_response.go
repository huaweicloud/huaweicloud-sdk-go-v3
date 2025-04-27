package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublicZoneLinesResponse Response Object
type ListPublicZoneLinesResponse struct {
	Links *PageLink `json:"links,omitempty"`

	// 查询公网域名的线路列表响应。
	Lines *[]PublicZoneLines `json:"lines,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPublicZoneLinesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicZoneLinesResponse struct{}"
	}

	return strings.Join([]string{"ListPublicZoneLinesResponse", string(data)}, " ")
}
