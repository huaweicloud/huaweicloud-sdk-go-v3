package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDrugLigandDifferenceResponse Response Object
type CheckDrugLigandDifferenceResponse struct {

	// 差异值
	Result         *float32 `json:"result,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o CheckDrugLigandDifferenceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDrugLigandDifferenceResponse struct{}"
	}

	return strings.Join([]string{"CheckDrugLigandDifferenceResponse", string(data)}, " ")
}
