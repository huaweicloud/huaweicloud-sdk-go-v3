package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerConfigsResponse Response Object
type ListServerConfigsResponse struct {
	Onpremises *OnPremisesConfig `json:"onpremises,omitempty"`

	Federations    *FederationConfig `json:"federations,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListServerConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListServerConfigsResponse", string(data)}, " ")
}
