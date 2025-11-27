package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GitRepository struct {
	Metadata *GitRepositoryMetaData `json:"metadata,omitempty"`

	Spec *GitRepositorySpec `json:"spec,omitempty"`

	Status *GitRepositoryStatus `json:"status,omitempty"`
}

func (o GitRepository) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GitRepository struct{}"
	}

	return strings.Join([]string{"GitRepository", string(data)}, " ")
}
