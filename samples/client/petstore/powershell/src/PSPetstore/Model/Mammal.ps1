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

Mammal<PSCustomObject>
#>
function ConvertFrom-PSJsonToMammal {
    [CmdletBinding()]
    Param (
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        $match = 0
        $matchType = $null
        $matchInstance = $null

        # try to match Pig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-PSJsonToPig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "Pig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'Pig' defined in oneOf (PSMammal). Proceeding to the next one if any."
        }

        # try to match Whale defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-PSJsonToWhale $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "Whale"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'Whale' defined in oneOf (PSMammal). Proceeding to the next one if any."
        }

        # try to match Zebra defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-PSJsonToZebra $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "Zebra"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'Zebra' defined in oneOf (PSMammal). Proceeding to the next one if any."
        }

        if ($match -gt 1) {
            throw "Error! The JSON payload matches more than one type defined in oneOf schemas ([Pig, Whale, Zebra]). JSON Payload: $($Json)"
        } elseif ($match -eq 1) {
            return [PSCustomObject]@{
                "ActualType" = ${matchType}
                "ActualInstance" = ${matchInstance}
                "OneOfSchemas" = @("Pig", "Whale", "Zebra")
            }
        } else {
            throw "Error! The JSON payload doesn't matches any type defined in oneOf schemas ([Pig, Whale, Zebra]). JSON Payload: $($Json)"
        }
    }
}

