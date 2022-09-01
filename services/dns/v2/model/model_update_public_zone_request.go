package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePublicZoneRequest struct {

	// 待修改zone的ID
	ZoneId string `json:"zone_id" xml:"zone_id"`

	Body *UpdatePublicZoneInfo `json:"body,omitempty" xml:"body"`
}

func (o UpdatePublicZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneRequest struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneRequest", string(data)}, " ")
}
