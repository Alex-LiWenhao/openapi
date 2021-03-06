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

.PARAMETER ArrayOfString
No description available.
.PARAMETER ArrayArrayOfInteger
No description available.
.PARAMETER ArrayArrayOfModel
No description available.
.OUTPUTS

ArrayTest<PSCustomObject>
#>

function Initialize-PSArrayTest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String[]]
        ${ArrayOfString},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [Int64[][]]
        ${ArrayArrayOfInteger},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[][]]
        ${ArrayArrayOfModel}
    )

    Process {
        'Creating PSCustomObject: PSPetstore => PSArrayTest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "array_of_string" = ${ArrayOfString}
            "array_array_of_integer" = ${ArrayArrayOfInteger}
            "array_array_of_model" = ${ArrayArrayOfModel}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ArrayTest<PSCustomObject>

.DESCRIPTION

Convert from JSON to ArrayTest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ArrayTest<PSCustomObject>
#>
function ConvertFrom-PSJsonToArrayTest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSPetstore => PSArrayTest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in PSArrayTest
        $AllProperties = ("array_of_string", "array_array_of_integer", "array_array_of_model")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "array_of_string"))) { #optional property not found
            $ArrayOfString = $null
        } else {
            $ArrayOfString = $JsonParameters.PSobject.Properties["array_of_string"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "array_array_of_integer"))) { #optional property not found
            $ArrayArrayOfInteger = $null
        } else {
            $ArrayArrayOfInteger = $JsonParameters.PSobject.Properties["array_array_of_integer"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "array_array_of_model"))) { #optional property not found
            $ArrayArrayOfModel = $null
        } else {
            $ArrayArrayOfModel = $JsonParameters.PSobject.Properties["array_array_of_model"].value
        }

        $PSO = [PSCustomObject]@{
            "array_of_string" = ${ArrayOfString}
            "array_array_of_integer" = ${ArrayArrayOfInteger}
            "array_array_of_model" = ${ArrayArrayOfModel}
        }

        return $PSO
    }

}

