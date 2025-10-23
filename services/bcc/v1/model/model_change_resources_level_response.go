package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeResourcesLevelResponse Response Object
type ChangeResourcesLevelResponse struct {

	// 每个资源更改等级的结果
	Results        *[]ResourceChangeLevelItem `json:"results,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ChangeResourcesLevelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeResourcesLevelResponse struct{}"
	}

	return strings.Join([]string{"ChangeResourcesLevelResponse", string(data)}, " ")
}
