package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExitEntryPermitEndorsementInfo struct {

	// 签注种类。
	EndorsementType *string `json:"endorsement_type,omitempty" xml:"endorsement_type"`

	// 签注往返有效次数。
	ValidRoundTrips *string `json:"valid_round_trips,omitempty" xml:"valid_round_trips"`

	// 签注有效期。
	EndorsementValidPeriod *string `json:"endorsement_valid_period,omitempty" xml:"endorsement_valid_period"`

	// 签注备注。
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// 签注签发信息。
	IssueInfo *string `json:"issue_info,omitempty" xml:"issue_info"`
}

func (o ExitEntryPermitEndorsementInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExitEntryPermitEndorsementInfo struct{}"
	}

	return strings.Join([]string{"ExitEntryPermitEndorsementInfo", string(data)}, " ")
}
