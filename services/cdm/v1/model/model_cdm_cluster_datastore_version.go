package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CdmClusterDatastoreVersion 版本。
type CdmClusterDatastoreVersion struct {

	// 版本ID。
	Id *string `json:"id,omitempty"`

	// 版本名称。
	Name *string `json:"name,omitempty"`

	// 规格信息。
	Flavors *[]CdmClusterFlavor `json:"flavors,omitempty"`
}

func (o CdmClusterDatastoreVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmClusterDatastoreVersion struct{}"
	}

	return strings.Join([]string{"CdmClusterDatastoreVersion", string(data)}, " ")
}
