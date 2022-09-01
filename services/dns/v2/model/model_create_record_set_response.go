package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateRecordSetResponse struct {
	Id *string `json:"id,omitempty" xml:"id"`

	Name *string `json:"name,omitempty" xml:"name"`

	Description *string `json:"description,omitempty" xml:"description"`

	ZoneId *string `json:"zone_id,omitempty" xml:"zone_id"`

	ZoneName *string `json:"zone_name,omitempty" xml:"zone_name"`

	Type *string `json:"type,omitempty" xml:"type"`

	Ttl *int32 `json:"ttl,omitempty" xml:"ttl"`

	Records *[]string `json:"records,omitempty" xml:"records"`

	CreateAt *string `json:"create_at,omitempty" xml:"create_at"`

	UpdateAt *string `json:"update_at,omitempty" xml:"update_at"`

	Status *string `json:"status,omitempty" xml:"status"`

	Default *bool `json:"default,omitempty" xml:"default"`

	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	Links          *PageLink `json:"links,omitempty" xml:"links"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateRecordSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetResponse struct{}"
	}

	return strings.Join([]string{"CreateRecordSetResponse", string(data)}, " ")
}
