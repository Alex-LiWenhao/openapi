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

.PARAMETER Json

JSON object

.OUTPUTS

Triangle<PSCustomObject>
#>
function ConvertFrom-PSJsonToTriangle {
    [CmdletBinding()]
    Param (
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        $match = 0
        $matchType = $null
        $matchInstance = $null

        # try to match EquilateralTriangle defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-PSJsonToEquilateralTriangle $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "EquilateralTriangle"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'EquilateralTriangle' defined in oneOf (PSTriangle). Proceeding to the next one if any."
        }

        # try to match IsoscelesTriangle defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-PSJsonToIsoscelesTriangle $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "IsoscelesTriangle"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'IsoscelesTriangle' defined in oneOf (PSTriangle). Proceeding to the next one if any."
        }

        # try to match ScaleneTriangle defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-PSJsonToScaleneTriangle $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "ScaleneTriangle"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'ScaleneTriangle' defined in oneOf (PSTriangle). Proceeding to the next one if any."
        }

        if ($match -gt 1) {
            throw "Error! The JSON payload matches more than one type defined in oneOf schemas ([EquilateralTriangle, IsoscelesTriangle, ScaleneTriangle]). JSON Payload: $($Json)"
        } elseif ($match -eq 1) {
            return [PSCustomObject]@{
                "ActualType" = ${matchType}
                "ActualInstance" = ${matchInstance}
                "OneOfSchemas" = @("EquilateralTriangle", "IsoscelesTriangle", "ScaleneTriangle")
            }
        } else {
            throw "Error! The JSON payload doesn't matches any type defined in oneOf schemas ([EquilateralTriangle, IsoscelesTriangle, ScaleneTriangle]). JSON Payload: $($Json)"
        }
    }
}

