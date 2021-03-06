package org.openapitools.server.model


/**
 * = An uploaded response =
 *
 * Describes the result of uploading an image resource
 *
 * @param code  for example: ''null''
 * @param `type`  for example: ''null''
 * @param message  for example: ''null''
*/
final case class ApiResponse (
  code: Option[Int] = None,
  `type`: Option[String] = None,
  message: Option[String] = None
)

