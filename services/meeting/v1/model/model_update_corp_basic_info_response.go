package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCorpBasicInfoResponse Response Object
type UpdateCorpBasicInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCorpBasicInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCorpBasicInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateCorpBasicInfoResponse", string(data)}, " ")
}
