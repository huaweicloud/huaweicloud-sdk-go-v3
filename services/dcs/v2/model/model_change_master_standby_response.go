package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeMasterStandbyResponse Response Object
type ChangeMasterStandbyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeMasterStandbyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeMasterStandbyResponse struct{}"
	}

	return strings.Join([]string{"ChangeMasterStandbyResponse", string(data)}, " ")
}
