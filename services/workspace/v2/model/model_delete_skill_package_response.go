package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSkillPackageResponse Response Object
type DeleteSkillPackageResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSkillPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSkillPackageResponse struct{}"
	}

	return strings.Join([]string{"DeleteSkillPackageResponse", string(data)}, " ")
}
