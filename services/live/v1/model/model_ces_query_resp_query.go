package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CesQueryRespQuery ces查询响应内容
type CesQueryRespQuery struct {
	MedialiveMpc *CesDimsItem `json:"medialive_mpc"`

	Pipeline *CesDimsItem `json:"pipeline"`

	Audio *CesDimsItem `json:"audio"`

	MedialiveCdn *CesDimsItem `json:"medialive_cdn"`

	MedialivePackage *CesDimsItem `json:"medialive_package"`

	MedialiveConnect *CesDimsItem `json:"medialive_connect"`

	MedialiveTailor *CesDimsItem `json:"medialive_tailor"`

	Region *CesDimsItem `json:"region"`
}

func (o CesQueryRespQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CesQueryRespQuery struct{}"
	}

	return strings.Join([]string{"CesQueryRespQuery", string(data)}, " ")
}
