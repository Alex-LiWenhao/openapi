#
# OpenAPI Petstore
# This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: "" \
# Version: 1.0.0
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

<#
.SYNOPSIS

No summary available.

.DESCRIPTION

No description available.

.PARAMETER PetType
No description available.
.OUTPUTS

ParentPet<PSCustomObject>
#>

function Initialize-PSParentPet {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${PetType}
    )

    Process {
        'Creating PSCustomObject: PSPetstore => PSParentPet' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $PetType) {
            throw "invalid value for 'PetType', 'PetType' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "pet_type" = ${PetType}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ParentPet<PSCustomObject>

.DESCRIPTION

Convert from JSON to ParentPet<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ParentPet<PSCustomObject>
#>
function ConvertFrom-PSJsonToParentPet {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSPetstore => PSParentPet' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in PSParentPet
        $AllProperties = ("pet_type")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'pet_type' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "pet_type"))) {
            throw "Error! JSON cannot be serialized due to the required property 'pet_type' missing."
        } else {
            $PetType = $JsonParameters.PSobject.Properties["pet_type"].value
        }

        $PSO = [PSCustomObject]@{
            "pet_type" = ${PetType}
        }

        return $PSO
    }

}

