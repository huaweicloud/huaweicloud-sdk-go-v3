package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAlertRspProcess struct {

	// The name, display only
	ProcessName *string `json:"process_name,omitempty"`

	// The name, display only
	ProcessPath *string `json:"process_path,omitempty"`

	// Id value
	ProcessPid *string `json:"process_pid,omitempty"`

	// Id value
	ProcessUid *string `json:"process_uid,omitempty"`

	// The name, display only
	ProcessCmdline *string `json:"process_cmdline,omitempty"`
}

func (o ShowAlertRspProcess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRspProcess struct{}"
	}

	return strings.Join([]string{"ShowAlertRspProcess", string(data)}, " ")
}
