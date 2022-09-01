package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ValidateSqlResponse struct {

	// SQL是否通过语法校验
	Valid          *bool `json:"valid,omitempty" xml:"valid"`
	HttpStatusCode int   `json:"-"`
}

func (o ValidateSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateSqlResponse struct{}"
	}

	return strings.Join([]string{"ValidateSqlResponse", string(data)}, " ")
}
