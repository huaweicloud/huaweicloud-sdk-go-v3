package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertItem struct {
	Metadata *MetaCert `json:"metadata,omitempty"`

	Spec *SpecCert `json:"spec,omitempty"`
}

func (o CertItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertItem struct{}"
	}

	return strings.Join([]string{"CertItem", string(data)}, " ")
}
