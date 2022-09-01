package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBlockchainStatusResponse struct {
	Bcs *Detail `json:"bcs,omitempty" xml:"bcs"`

	Eip *Detail `json:"eip,omitempty" xml:"eip"`

	Sfs *Detail `json:"sfs,omitempty" xml:"sfs"`

	Obs *Detail `json:"obs,omitempty" xml:"obs"`

	Kafka *Detail `json:"kafka,omitempty" xml:"kafka"`

	Cce            *ComCce `json:"cce,omitempty" xml:"cce"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBlockchainStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockchainStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowBlockchainStatusResponse", string(data)}, " ")
}
