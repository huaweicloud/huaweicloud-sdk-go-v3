package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBaremetalServerTagsResponse Response Object
type ShowBaremetalServerTagsResponse struct {

	//
	Tags           *[]BaremetalServerTag `json:"tags,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowBaremetalServerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaremetalServerTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowBaremetalServerTagsResponse", string(data)}, " ")
}
