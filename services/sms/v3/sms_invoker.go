package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sms/v3/model"
)

type CheckNetAclInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckNetAclInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckNetAclInvoker) Invoke() (*model.CheckNetAclResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckNetAclResponse), nil
	}
}

type CollectLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CollectLogInvoker) Invoke() (*model.CollectLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectLogResponse), nil
	}
}

type CreateMigprojectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMigprojectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMigprojectInvoker) Invoke() (*model.CreateMigprojectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMigprojectResponse), nil
	}
}

type CreatePrivacyAgreementsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrivacyAgreementsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePrivacyAgreementsInvoker) Invoke() (*model.CreatePrivacyAgreementsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivacyAgreementsResponse), nil
	}
}

type CreateTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTaskInvoker) Invoke() (*model.CreateTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTaskResponse), nil
	}
}

type CreateTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTemplateInvoker) Invoke() (*model.CreateTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTemplateResponse), nil
	}
}

type DeleteMigprojectInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMigprojectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMigprojectInvoker) Invoke() (*model.DeleteMigprojectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMigprojectResponse), nil
	}
}

type DeleteServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServerInvoker) Invoke() (*model.DeleteServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServerResponse), nil
	}
}

type DeleteServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServersInvoker) Invoke() (*model.DeleteServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServersResponse), nil
	}
}

type DeleteTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTaskInvoker) Invoke() (*model.DeleteTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTaskResponse), nil
	}
}

type DeleteTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTasksInvoker) Invoke() (*model.DeleteTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTasksResponse), nil
	}
}

type DeleteTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTemplateInvoker) Invoke() (*model.DeleteTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplateResponse), nil
	}
}

type DeleteTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTemplatesInvoker) Invoke() (*model.DeleteTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplatesResponse), nil
	}
}

type ListErrorServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListErrorServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListErrorServersInvoker) Invoke() (*model.ListErrorServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListErrorServersResponse), nil
	}
}

type ListMigprojectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMigprojectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMigprojectsInvoker) Invoke() (*model.ListMigprojectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMigprojectsResponse), nil
	}
}

type ListServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServersInvoker) Invoke() (*model.ListServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServersResponse), nil
	}
}

type ListTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTasksInvoker) Invoke() (*model.ListTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTasksResponse), nil
	}
}

type ListTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTemplatesInvoker) Invoke() (*model.ListTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplatesResponse), nil
	}
}

type RegisterServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterServerInvoker) Invoke() (*model.RegisterServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterServerResponse), nil
	}
}

type ShowCertKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCertKeyInvoker) Invoke() (*model.ShowCertKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertKeyResponse), nil
	}
}

type ShowCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCommandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCommandInvoker) Invoke() (*model.ShowCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCommandResponse), nil
	}
}

type ShowConfigSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConfigSettingInvoker) Invoke() (*model.ShowConfigSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigSettingResponse), nil
	}
}

type ShowConsistencyResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConsistencyResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConsistencyResultInvoker) Invoke() (*model.ShowConsistencyResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConsistencyResultResponse), nil
	}
}

type ShowMigprojectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMigprojectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMigprojectInvoker) Invoke() (*model.ShowMigprojectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMigprojectResponse), nil
	}
}

type ShowOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOverviewInvoker) Invoke() (*model.ShowOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOverviewResponse), nil
	}
}

type ShowPassphraseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPassphraseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPassphraseInvoker) Invoke() (*model.ShowPassphraseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPassphraseResponse), nil
	}
}

type ShowPrivacyAgreementsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivacyAgreementsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPrivacyAgreementsInvoker) Invoke() (*model.ShowPrivacyAgreementsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivacyAgreementsResponse), nil
	}
}

type ShowServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerInvoker) Invoke() (*model.ShowServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerResponse), nil
	}
}

type ShowSha256Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSha256Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSha256Invoker) Invoke() (*model.ShowSha256Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSha256Response), nil
	}
}

type ShowTargetPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTargetPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTargetPasswordInvoker) Invoke() (*model.ShowTargetPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTargetPasswordResponse), nil
	}
}

type ShowTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskInvoker) Invoke() (*model.ShowTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskResponse), nil
	}
}

type ShowTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTemplateInvoker) Invoke() (*model.ShowTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateResponse), nil
	}
}

type ShowsSpeedLimitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowsSpeedLimitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowsSpeedLimitsInvoker) Invoke() (*model.ShowsSpeedLimitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowsSpeedLimitsResponse), nil
	}
}

type UnlockTargetEcsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnlockTargetEcsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnlockTargetEcsInvoker) Invoke() (*model.UnlockTargetEcsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnlockTargetEcsResponse), nil
	}
}

type UpdateCommandResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCommandResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCommandResultInvoker) Invoke() (*model.UpdateCommandResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCommandResultResponse), nil
	}
}

type UpdateConsistencyResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConsistencyResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateConsistencyResultInvoker) Invoke() (*model.UpdateConsistencyResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConsistencyResultResponse), nil
	}
}

type UpdateCopyStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCopyStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCopyStateInvoker) Invoke() (*model.UpdateCopyStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCopyStateResponse), nil
	}
}

type UpdateDefaultMigprojectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDefaultMigprojectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDefaultMigprojectInvoker) Invoke() (*model.UpdateDefaultMigprojectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDefaultMigprojectResponse), nil
	}
}

type UpdateDiskInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDiskInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDiskInfoInvoker) Invoke() (*model.UpdateDiskInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDiskInfoResponse), nil
	}
}

type UpdateMigprojectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMigprojectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMigprojectInvoker) Invoke() (*model.UpdateMigprojectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMigprojectResponse), nil
	}
}

type UpdateNetworkCheckInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNetworkCheckInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNetworkCheckInfoInvoker) Invoke() (*model.UpdateNetworkCheckInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNetworkCheckInfoResponse), nil
	}
}

type UpdateServerNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServerNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateServerNameInvoker) Invoke() (*model.UpdateServerNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServerNameResponse), nil
	}
}

type UpdateSpeedInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSpeedInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSpeedInvoker) Invoke() (*model.UpdateSpeedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSpeedResponse), nil
	}
}

type UpdateTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTaskInvoker) Invoke() (*model.UpdateTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaskResponse), nil
	}
}

type UpdateTaskSpeedInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaskSpeedInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTaskSpeedInvoker) Invoke() (*model.UpdateTaskSpeedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaskSpeedResponse), nil
	}
}

type UpdateTaskStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaskStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTaskStatusInvoker) Invoke() (*model.UpdateTaskStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaskStatusResponse), nil
	}
}

type UpdateTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTemplateInvoker) Invoke() (*model.UpdateTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTemplateResponse), nil
	}
}

type UploadSpecialConfigurationSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadSpecialConfigurationSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadSpecialConfigurationSettingInvoker) Invoke() (*model.UploadSpecialConfigurationSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadSpecialConfigurationSettingResponse), nil
	}
}

type ShowConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConfigInvoker) Invoke() (*model.ShowConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigResponse), nil
	}
}

type ListApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiVersionInvoker) Invoke() (*model.ListApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionResponse), nil
	}
}

type ShowApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApiVersionInvoker) Invoke() (*model.ShowApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiVersionResponse), nil
	}
}
