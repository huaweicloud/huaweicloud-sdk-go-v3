package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Kustomization struct {
	Metadata *KustomizationMetaData `json:"metadata,omitempty"`

	Spec *KustomizationSpec `json:"spec,omitempty"`

	Status *KustomizationStatus `json:"status,omitempty"`
}

func (o Kustomization) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Kustomization struct{}"
	}

	return strings.Join([]string{"Kustomization", string(data)}, " ")
}
