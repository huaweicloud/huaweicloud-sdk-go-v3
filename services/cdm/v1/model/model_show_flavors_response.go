package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlavorsResponse Response Object
type ShowFlavorsResponse struct {

	// 服务ID，用于区分不同服务。
	Id *string `json:"id,omitempty"`

	// db名称，一般为cdm。
	Dbname *string `json:"dbname,omitempty"`

	// 版本信息列表。
	Versions       *[]CdmClusterDatastoreVersion `json:"versions,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ShowFlavorsResponse", string(data)}, " ")
}
