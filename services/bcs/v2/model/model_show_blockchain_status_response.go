package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBlockchainStatusResponse Response Object
type ShowBlockchainStatusResponse struct {
	Bcs *Detail `json:"bcs,omitempty"`

	Eip *Detail `json:"eip,omitempty"`

	Sfs *Detail `json:"sfs,omitempty"`

	Obs *Detail `json:"obs,omitempty"`

	Kafka *Detail `json:"kafka,omitempty"`

	Cce            *ComCce `json:"cce,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBlockchainStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockchainStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowBlockchainStatusResponse", string(data)}, " ")
}
