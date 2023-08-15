package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryArchiveResponse Response Object
type ShowRepositoryArchiveResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ShowRepositoryArchiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryArchiveResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryArchiveResponse", string(data)}, " ")
}
