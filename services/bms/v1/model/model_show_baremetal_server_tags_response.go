package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBaremetalServerTagsResponse struct {

	//
	Tags           *[]BaremetalServerTag `json:"tags,omitempty" xml:"tags"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowBaremetalServerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaremetalServerTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowBaremetalServerTagsResponse", string(data)}, " ")
}
