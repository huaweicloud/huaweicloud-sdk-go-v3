package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CandidatesUser struct {

	// 用户名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 用户ID
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o CandidatesUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CandidatesUser struct{}"
	}

	return strings.Join([]string{"CandidatesUser", string(data)}, " ")
}
