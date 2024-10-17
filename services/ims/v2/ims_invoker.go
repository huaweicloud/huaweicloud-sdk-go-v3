package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2/model"
)

type AddImageTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddImageTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddImageTagInvoker) Invoke() (*model.AddImageTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddImageTagResponse), nil
	}
}

type BatchAddMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddMembersInvoker) Invoke() (*model.BatchAddMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddMembersResponse), nil
	}
}

type BatchAddOrDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddOrDeleteTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddOrDeleteTagsInvoker) Invoke() (*model.BatchAddOrDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddOrDeleteTagsResponse), nil
	}
}

type BatchDeleteMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteMembersInvoker) Invoke() (*model.BatchDeleteMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteMembersResponse), nil
	}
}

type BatchUpdateMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateMembersInvoker) Invoke() (*model.BatchUpdateMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateMembersResponse), nil
	}
}

type CopyImageCrossRegionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyImageCrossRegionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CopyImageCrossRegionInvoker) Invoke() (*model.CopyImageCrossRegionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyImageCrossRegionResponse), nil
	}
}

type CopyImageInRegionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyImageInRegionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CopyImageInRegionInvoker) Invoke() (*model.CopyImageInRegionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyImageInRegionResponse), nil
	}
}

type CreateDataImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDataImageInvoker) Invoke() (*model.CreateDataImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataImageResponse), nil
	}
}

type CreateImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateImageInvoker) Invoke() (*model.CreateImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageResponse), nil
	}
}

type CreateOrUpdateTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrUpdateTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrUpdateTagsInvoker) Invoke() (*model.CreateOrUpdateTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrUpdateTagsResponse), nil
	}
}

type CreateWholeImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWholeImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWholeImageInvoker) Invoke() (*model.CreateWholeImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWholeImageResponse), nil
	}
}

type DeleteImageTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteImageTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteImageTagInvoker) Invoke() (*model.DeleteImageTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteImageTagResponse), nil
	}
}

type ExportImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportImageInvoker) Invoke() (*model.ExportImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportImageResponse), nil
	}
}

type ImportImageQuickInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportImageQuickInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportImageQuickInvoker) Invoke() (*model.ImportImageQuickResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportImageQuickResponse), nil
	}
}

type ListImageByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageByTagsInvoker) Invoke() (*model.ListImageByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageByTagsResponse), nil
	}
}

type ListImageTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageTagsInvoker) Invoke() (*model.ListImageTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageTagsResponse), nil
	}
}

type ListImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImagesInvoker) Invoke() (*model.ListImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImagesResponse), nil
	}
}

type ListImagesTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImagesTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImagesTagsInvoker) Invoke() (*model.ListImagesTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImagesTagsResponse), nil
	}
}

type ListOsVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOsVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOsVersionsInvoker) Invoke() (*model.ListOsVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOsVersionsResponse), nil
	}
}

type ListTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagsInvoker) Invoke() (*model.ListTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsResponse), nil
	}
}

type RegisterImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterImageInvoker) Invoke() (*model.RegisterImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterImageResponse), nil
	}
}

type ShowImageQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageQuotaInvoker) Invoke() (*model.ShowImageQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageQuotaResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type ShowJobProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobProgressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobProgressInvoker) Invoke() (*model.ShowJobProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobProgressResponse), nil
	}
}

type UpdateImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateImageInvoker) Invoke() (*model.UpdateImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateImageResponse), nil
	}
}

type ListVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVersionsInvoker) Invoke() (*model.ListVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVersionsResponse), nil
	}
}

type ShowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVersionInvoker) Invoke() (*model.ShowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVersionResponse), nil
	}
}

type GlanceAddImageMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceAddImageMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceAddImageMemberInvoker) Invoke() (*model.GlanceAddImageMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceAddImageMemberResponse), nil
	}
}

type GlanceCreateImageMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceCreateImageMetadataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceCreateImageMetadataInvoker) Invoke() (*model.GlanceCreateImageMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceCreateImageMetadataResponse), nil
	}
}

type GlanceCreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceCreateTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceCreateTagInvoker) Invoke() (*model.GlanceCreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceCreateTagResponse), nil
	}
}

type GlanceDeleteImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceDeleteImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceDeleteImageInvoker) Invoke() (*model.GlanceDeleteImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceDeleteImageResponse), nil
	}
}

type GlanceDeleteImageMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceDeleteImageMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceDeleteImageMemberInvoker) Invoke() (*model.GlanceDeleteImageMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceDeleteImageMemberResponse), nil
	}
}

type GlanceDeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceDeleteTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceDeleteTagInvoker) Invoke() (*model.GlanceDeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceDeleteTagResponse), nil
	}
}

type GlanceListImageMemberSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceListImageMemberSchemasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceListImageMemberSchemasInvoker) Invoke() (*model.GlanceListImageMemberSchemasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceListImageMemberSchemasResponse), nil
	}
}

type GlanceListImageMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceListImageMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceListImageMembersInvoker) Invoke() (*model.GlanceListImageMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceListImageMembersResponse), nil
	}
}

type GlanceListImageSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceListImageSchemasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceListImageSchemasInvoker) Invoke() (*model.GlanceListImageSchemasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceListImageSchemasResponse), nil
	}
}

type GlanceListImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceListImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceListImagesInvoker) Invoke() (*model.GlanceListImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceListImagesResponse), nil
	}
}

type GlanceShowImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceShowImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceShowImageInvoker) Invoke() (*model.GlanceShowImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceShowImageResponse), nil
	}
}

type GlanceShowImageMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceShowImageMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceShowImageMemberInvoker) Invoke() (*model.GlanceShowImageMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceShowImageMemberResponse), nil
	}
}

type GlanceShowImageMemberSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceShowImageMemberSchemasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceShowImageMemberSchemasInvoker) Invoke() (*model.GlanceShowImageMemberSchemasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceShowImageMemberSchemasResponse), nil
	}
}

type GlanceShowImageSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceShowImageSchemasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceShowImageSchemasInvoker) Invoke() (*model.GlanceShowImageSchemasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceShowImageSchemasResponse), nil
	}
}

type GlanceUpdateImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceUpdateImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceUpdateImageInvoker) Invoke() (*model.GlanceUpdateImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceUpdateImageResponse), nil
	}
}

type GlanceUpdateImageMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *GlanceUpdateImageMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GlanceUpdateImageMemberInvoker) Invoke() (*model.GlanceUpdateImageMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GlanceUpdateImageMemberResponse), nil
	}
}
