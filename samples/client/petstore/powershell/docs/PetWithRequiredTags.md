# PetWithRequiredTags
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Category** | [**Category**](Category.md) |  | [optional] 
**Name** | **String** |  | 
**PhotoUrls** | **String[]** |  | 
**Tags** | [**Tag[]**](Tag.md) |  | 
**Status** | **String** | pet status in the store | [optional] 

## Examples

- Prepare the resource
```powershell
$PetWithRequiredTags = Initialize-PSPetstorePetWithRequiredTags  -Id null `
 -Category null `
 -Name doggie `
 -PhotoUrls null `
 -Tags null `
 -Status null
```

- Convert the resource to JSON
```powershell
$PetWithRequiredTags | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

