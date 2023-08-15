package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExitEntryPermitEndorsementInfo struct {

	// 签注种类。
	EndorsementType *string `json:"endorsement_type,omitempty"`

	// 签注往返有效次数。
	ValidRoundTrips *string `json:"valid_round_trips,omitempty"`

	// 签注有效期。
	EndorsementValidPeriod *string `json:"endorsement_valid_period,omitempty"`

	// 签注备注。
	Remark *string `json:"remark,omitempty"`

	// 签注签发信息。
	IssueInfo *string `json:"issue_info,omitempty"`
}

func (o ExitEntryPermitEndorsementInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExitEntryPermitEndorsementInfo struct{}"
	}

	return strings.Join([]string{"ExitEntryPermitEndorsementInfo", string(data)}, " ")
}
