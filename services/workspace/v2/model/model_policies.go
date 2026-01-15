package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Policies 策略类。
type Policies struct {
	Peripherals *PoliciesPeripherals `json:"peripherals,omitempty"`

	Audio *PoliciesAudio `json:"audio,omitempty"`

	Client *PoliciesClient `json:"client,omitempty"`

	Display *PoliciesDisplay `json:"display,omitempty"`

	FileAndClipboard *PoliciesFileAndClipboard `json:"file_and_clipboard,omitempty"`

	Session *Session `json:"session,omitempty"`

	VirtualChannel *VirtualChannel `json:"virtual_channel,omitempty"`

	Watermark *Watermark `json:"watermark,omitempty"`

	KeyboardMouse *PoliciesKeyboardMouse `json:"keyboard_mouse,omitempty"`

	Seamless *PoliciesSeamless `json:"seamless,omitempty"`

	PersonalizedDataMgmt *PoliciesPersonalizedDataMgmt `json:"personalized_data_mgmt,omitempty"`

	Custom *PoliciesCustom `json:"custom,omitempty"`

	RecordAudit *PoliciesRecordAudit `json:"record_audit,omitempty"`

	CloudStorage *PoliciesCloudStorage `json:"cloud_storage,omitempty"`

	UserProfile *PoliciesUserProfile `json:"user_profile,omitempty"`
}

func (o Policies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Policies struct{}"
	}

	return strings.Join([]string{"Policies", string(data)}, " ")
}
