package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenter/v1/model"
)

type CreateAccountAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccountAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAccountAssignmentInvoker) Invoke() (*model.CreateAccountAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccountAssignmentResponse), nil
	}
}

type DeleteAccountAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAccountAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAccountAssignmentInvoker) Invoke() (*model.DeleteAccountAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAccountAssignmentResponse), nil
	}
}

type DescribeAccountAssignmentCreationStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeAccountAssignmentCreationStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeAccountAssignmentCreationStatusInvoker) Invoke() (*model.DescribeAccountAssignmentCreationStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeAccountAssignmentCreationStatusResponse), nil
	}
}

type DescribeAccountAssignmentDeletionStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeAccountAssignmentDeletionStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeAccountAssignmentDeletionStatusInvoker) Invoke() (*model.DescribeAccountAssignmentDeletionStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeAccountAssignmentDeletionStatusResponse), nil
	}
}

type DisassociateProfileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateProfileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateProfileInvoker) Invoke() (*model.DisassociateProfileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateProfileResponse), nil
	}
}

type ListAccountAssignmentCreationStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountAssignmentCreationStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountAssignmentCreationStatusInvoker) Invoke() (*model.ListAccountAssignmentCreationStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountAssignmentCreationStatusResponse), nil
	}
}

type ListAccountAssignmentDeletionStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountAssignmentDeletionStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountAssignmentDeletionStatusInvoker) Invoke() (*model.ListAccountAssignmentDeletionStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountAssignmentDeletionStatusResponse), nil
	}
}

type ListAccountAssignmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountAssignmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountAssignmentsInvoker) Invoke() (*model.ListAccountAssignmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountAssignmentsResponse), nil
	}
}

type ListAccountAssignmentsForPrincipalInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountAssignmentsForPrincipalInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountAssignmentsForPrincipalInvoker) Invoke() (*model.ListAccountAssignmentsForPrincipalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountAssignmentsForPrincipalResponse), nil
	}
}

type CreateApplicationInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApplicationInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApplicationInstanceInvoker) Invoke() (*model.CreateApplicationInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApplicationInstanceResponse), nil
	}
}

type DeleteApplicationInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApplicationInstanceInvoker) Invoke() (*model.DeleteApplicationInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationInstanceResponse), nil
	}
}

type DeleteProfileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProfileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProfileInvoker) Invoke() (*model.DeleteProfileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProfileResponse), nil
	}
}

type DescribeApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeApplicationInvoker) Invoke() (*model.DescribeApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeApplicationResponse), nil
	}
}

type DescribeApplicationProviderInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeApplicationProviderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeApplicationProviderInvoker) Invoke() (*model.DescribeApplicationProviderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeApplicationProviderResponse), nil
	}
}

type GetApplicationAssignmentConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetApplicationAssignmentConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetApplicationAssignmentConfigurationInvoker) Invoke() (*model.GetApplicationAssignmentConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetApplicationAssignmentConfigurationResponse), nil
	}
}

type GetApplicationInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetApplicationInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetApplicationInstanceInvoker) Invoke() (*model.GetApplicationInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetApplicationInstanceResponse), nil
	}
}

type ImportApplicationInstanceServiceProviderMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportApplicationInstanceServiceProviderMetadataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportApplicationInstanceServiceProviderMetadataInvoker) Invoke() (*model.ImportApplicationInstanceServiceProviderMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportApplicationInstanceServiceProviderMetadataResponse), nil
	}
}

type ListApplicationInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationInstancesInvoker) Invoke() (*model.ListApplicationInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationInstancesResponse), nil
	}
}

type ListApplicationProvidersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationProvidersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationProvidersInvoker) Invoke() (*model.ListApplicationProvidersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationProvidersResponse), nil
	}
}

type ListApplicationTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationTemplatesInvoker) Invoke() (*model.ListApplicationTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationTemplatesResponse), nil
	}
}

type ListApplicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationsInvoker) Invoke() (*model.ListApplicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationsResponse), nil
	}
}

type ListCatalogApplicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCatalogApplicationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCatalogApplicationsInvoker) Invoke() (*model.ListCatalogApplicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCatalogApplicationsResponse), nil
	}
}

type ListProfilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProfilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProfilesInvoker) Invoke() (*model.ListProfilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProfilesResponse), nil
	}
}

type UpdateApplicationInstanceDisplayDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApplicationInstanceDisplayDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApplicationInstanceDisplayDataInvoker) Invoke() (*model.UpdateApplicationInstanceDisplayDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationInstanceDisplayDataResponse), nil
	}
}

type UpdateApplicationInstanceResponseConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApplicationInstanceResponseConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApplicationInstanceResponseConfigurationInvoker) Invoke() (*model.UpdateApplicationInstanceResponseConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationInstanceResponseConfigurationResponse), nil
	}
}

type UpdateApplicationInstanceResponseSchemaConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApplicationInstanceResponseSchemaConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApplicationInstanceResponseSchemaConfigurationInvoker) Invoke() (*model.UpdateApplicationInstanceResponseSchemaConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationInstanceResponseSchemaConfigurationResponse), nil
	}
}

type UpdateApplicationInstanceSecurityConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApplicationInstanceSecurityConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApplicationInstanceSecurityConfigurationInvoker) Invoke() (*model.UpdateApplicationInstanceSecurityConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationInstanceSecurityConfigurationResponse), nil
	}
}

type UpdateApplicationInstanceServiceProviderConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApplicationInstanceServiceProviderConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApplicationInstanceServiceProviderConfigurationInvoker) Invoke() (*model.UpdateApplicationInstanceServiceProviderConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationInstanceServiceProviderConfigurationResponse), nil
	}
}

type UpdateApplicationInstanceStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApplicationInstanceStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApplicationInstanceStatusInvoker) Invoke() (*model.UpdateApplicationInstanceStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationInstanceStatusResponse), nil
	}
}

type CreateApplicationAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApplicationAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApplicationAssignmentInvoker) Invoke() (*model.CreateApplicationAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApplicationAssignmentResponse), nil
	}
}

type DeleteApplicationAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApplicationAssignmentInvoker) Invoke() (*model.DeleteApplicationAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationAssignmentResponse), nil
	}
}

type ListApplicationAssignmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationAssignmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationAssignmentsInvoker) Invoke() (*model.ListApplicationAssignmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationAssignmentsResponse), nil
	}
}

type ListApplicationAssignmentsForPrincipalInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationAssignmentsForPrincipalInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationAssignmentsForPrincipalInvoker) Invoke() (*model.ListApplicationAssignmentsForPrincipalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationAssignmentsForPrincipalResponse), nil
	}
}

type CreateApplicationInstanceCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApplicationInstanceCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApplicationInstanceCertificateInvoker) Invoke() (*model.CreateApplicationInstanceCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApplicationInstanceCertificateResponse), nil
	}
}

type DeleteApplicationInstanceCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationInstanceCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApplicationInstanceCertificateInvoker) Invoke() (*model.DeleteApplicationInstanceCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationInstanceCertificateResponse), nil
	}
}

type ListApplicationInstanceCertificatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationInstanceCertificatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationInstanceCertificatesInvoker) Invoke() (*model.ListApplicationInstanceCertificatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationInstanceCertificatesResponse), nil
	}
}

type UpdateApplicationInstanceActiveCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApplicationInstanceActiveCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApplicationInstanceActiveCertificateInvoker) Invoke() (*model.UpdateApplicationInstanceActiveCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationInstanceActiveCertificateResponse), nil
	}
}

type GetSsoConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetSsoConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetSsoConfigurationInvoker) Invoke() (*model.GetSsoConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetSsoConfigurationResponse), nil
	}
}

type UpdateSsoConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSsoConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSsoConfigurationInvoker) Invoke() (*model.UpdateSsoConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSsoConfigurationResponse), nil
	}
}

type CreateAliasInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAliasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAliasInvoker) Invoke() (*model.CreateAliasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAliasResponse), nil
	}
}

type DeleteIdentityCenterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIdentityCenterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteIdentityCenterInvoker) Invoke() (*model.DeleteIdentityCenterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIdentityCenterResponse), nil
	}
}

type DescribeRegisteredRegionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeRegisteredRegionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeRegisteredRegionsInvoker) Invoke() (*model.DescribeRegisteredRegionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeRegisteredRegionsResponse), nil
	}
}

type GetHaConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetHaConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetHaConfigurationInvoker) Invoke() (*model.GetHaConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetHaConfigurationResponse), nil
	}
}

type GetIdentityCenterServiceStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetIdentityCenterServiceStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetIdentityCenterServiceStatusInvoker) Invoke() (*model.GetIdentityCenterServiceStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetIdentityCenterServiceStatusResponse), nil
	}
}

type ListIdentityStoreAssociationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIdentityStoreAssociationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIdentityStoreAssociationInvoker) Invoke() (*model.ListIdentityStoreAssociationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIdentityStoreAssociationResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type RegisterRegionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterRegionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterRegionInvoker) Invoke() (*model.RegisterRegionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterRegionResponse), nil
	}
}

type StartIdentityCenterInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartIdentityCenterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartIdentityCenterInvoker) Invoke() (*model.StartIdentityCenterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartIdentityCenterResponse), nil
	}
}

type UpdateHaConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHaConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHaConfigurationInvoker) Invoke() (*model.UpdateHaConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHaConfigurationResponse), nil
	}
}

type CreateInstanceAccessControlAttributeConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceAccessControlAttributeConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceAccessControlAttributeConfigurationInvoker) Invoke() (*model.CreateInstanceAccessControlAttributeConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceAccessControlAttributeConfigurationResponse), nil
	}
}

type DeleteInstanceAccessControlAttributeConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceAccessControlAttributeConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceAccessControlAttributeConfigurationInvoker) Invoke() (*model.DeleteInstanceAccessControlAttributeConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceAccessControlAttributeConfigurationResponse), nil
	}
}

type DescribeInstanceAccessControlAttributeConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeInstanceAccessControlAttributeConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeInstanceAccessControlAttributeConfigurationInvoker) Invoke() (*model.DescribeInstanceAccessControlAttributeConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeInstanceAccessControlAttributeConfigurationResponse), nil
	}
}

type UpdateInstanceAccessControlAttributeConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceAccessControlAttributeConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceAccessControlAttributeConfigurationInvoker) Invoke() (*model.UpdateInstanceAccessControlAttributeConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceAccessControlAttributeConfigurationResponse), nil
	}
}

type GetMfaDeviceManagementForIdentityStoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetMfaDeviceManagementForIdentityStoreInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetMfaDeviceManagementForIdentityStoreInvoker) Invoke() (*model.GetMfaDeviceManagementForIdentityStoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetMfaDeviceManagementForIdentityStoreResponse), nil
	}
}

type PutMfaDeviceManagementForIdentityStoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutMfaDeviceManagementForIdentityStoreInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PutMfaDeviceManagementForIdentityStoreInvoker) Invoke() (*model.PutMfaDeviceManagementForIdentityStoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutMfaDeviceManagementForIdentityStoreResponse), nil
	}
}

type AttachManagedPolicyToPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachManagedPolicyToPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachManagedPolicyToPermissionSetInvoker) Invoke() (*model.AttachManagedPolicyToPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachManagedPolicyToPermissionSetResponse), nil
	}
}

type AttachManagedRoleToPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachManagedRoleToPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachManagedRoleToPermissionSetInvoker) Invoke() (*model.AttachManagedRoleToPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachManagedRoleToPermissionSetResponse), nil
	}
}

type CreatePermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePermissionSetInvoker) Invoke() (*model.CreatePermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePermissionSetResponse), nil
	}
}

type DeleteCustomPolicyFromPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCustomPolicyFromPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCustomPolicyFromPermissionSetInvoker) Invoke() (*model.DeleteCustomPolicyFromPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCustomPolicyFromPermissionSetResponse), nil
	}
}

type DeleteCustomRoleFromPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCustomRoleFromPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCustomRoleFromPermissionSetInvoker) Invoke() (*model.DeleteCustomRoleFromPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCustomRoleFromPermissionSetResponse), nil
	}
}

type DeletePermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePermissionSetInvoker) Invoke() (*model.DeletePermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePermissionSetResponse), nil
	}
}

type DescribePermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribePermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribePermissionSetInvoker) Invoke() (*model.DescribePermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribePermissionSetResponse), nil
	}
}

type DescribePermissionSetProvisioningStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribePermissionSetProvisioningStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribePermissionSetProvisioningStatusInvoker) Invoke() (*model.DescribePermissionSetProvisioningStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribePermissionSetProvisioningStatusResponse), nil
	}
}

type DetachManagedPolicyFromPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachManagedPolicyFromPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachManagedPolicyFromPermissionSetInvoker) Invoke() (*model.DetachManagedPolicyFromPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachManagedPolicyFromPermissionSetResponse), nil
	}
}

type DetachManagedRoleFromPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachManagedRoleFromPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachManagedRoleFromPermissionSetInvoker) Invoke() (*model.DetachManagedRoleFromPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachManagedRoleFromPermissionSetResponse), nil
	}
}

type GetCustomPolicyForPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetCustomPolicyForPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetCustomPolicyForPermissionSetInvoker) Invoke() (*model.GetCustomPolicyForPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetCustomPolicyForPermissionSetResponse), nil
	}
}

type GetCustomRoleForPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetCustomRoleForPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetCustomRoleForPermissionSetInvoker) Invoke() (*model.GetCustomRoleForPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetCustomRoleForPermissionSetResponse), nil
	}
}

type GetPermissionSetSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetPermissionSetSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetPermissionSetSummaryInvoker) Invoke() (*model.GetPermissionSetSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetPermissionSetSummaryResponse), nil
	}
}

type ListAccountsForProvisionedPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountsForProvisionedPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountsForProvisionedPermissionSetInvoker) Invoke() (*model.ListAccountsForProvisionedPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountsForProvisionedPermissionSetResponse), nil
	}
}

type ListManagedPoliciesInPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListManagedPoliciesInPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListManagedPoliciesInPermissionSetInvoker) Invoke() (*model.ListManagedPoliciesInPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListManagedPoliciesInPermissionSetResponse), nil
	}
}

type ListManagedRolesInPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListManagedRolesInPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListManagedRolesInPermissionSetInvoker) Invoke() (*model.ListManagedRolesInPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListManagedRolesInPermissionSetResponse), nil
	}
}

type ListPermissionSetProvisioningStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPermissionSetProvisioningStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPermissionSetProvisioningStatusInvoker) Invoke() (*model.ListPermissionSetProvisioningStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPermissionSetProvisioningStatusResponse), nil
	}
}

type ListPermissionSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPermissionSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPermissionSetsInvoker) Invoke() (*model.ListPermissionSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPermissionSetsResponse), nil
	}
}

type ListPermissionSetsProvisionedToAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPermissionSetsProvisionedToAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPermissionSetsProvisionedToAccountInvoker) Invoke() (*model.ListPermissionSetsProvisionedToAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPermissionSetsProvisionedToAccountResponse), nil
	}
}

type ProvisionPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ProvisionPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ProvisionPermissionSetInvoker) Invoke() (*model.ProvisionPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProvisionPermissionSetResponse), nil
	}
}

type PutCustomPolicyToPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutCustomPolicyToPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PutCustomPolicyToPermissionSetInvoker) Invoke() (*model.PutCustomPolicyToPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutCustomPolicyToPermissionSetResponse), nil
	}
}

type PutCustomRoleToPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutCustomRoleToPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PutCustomRoleToPermissionSetInvoker) Invoke() (*model.PutCustomRoleToPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutCustomRoleToPermissionSetResponse), nil
	}
}

type UpdatePermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePermissionSetInvoker) Invoke() (*model.UpdatePermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePermissionSetResponse), nil
	}
}

type CreateTagResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTagResourceInvoker) Invoke() (*model.CreateTagResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagResourceResponse), nil
	}
}

type DeleteTagResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTagResourceInvoker) Invoke() (*model.DeleteTagResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResourceResponse), nil
	}
}

type ListTagResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagResourcesInvoker) Invoke() (*model.ListTagResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagResourcesResponse), nil
	}
}
