/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Net;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Azure.WebJobs;
using Microsoft.Azure.WebJobs.Extensions.Http;

namespace Org.OpenAPITools.Apis
{ 
    public partial class UserApi
    { 
        [FunctionName("UserApi_CreateUser")]
        public async Task<IActionResult> _CreateUser([HttpTrigger(AuthorizationLevel.Anonymous, "POST", Route = "/user")]HttpRequest req, ExecutionContext context)
        {
            var method = this.GetType().GetMethod("CreateUser");

            return method != null 
                ? (await ((Task<IActionResult>)method.Invoke(this, new object[] { req, context })).ConfigureAwait(false))
                : new StatusCodeResult((int)HttpStatusCode.NotImplemented);
        }

        [FunctionName("UserApi_CreateUsersWithArrayInput")]
        public async Task<IActionResult> _CreateUsersWithArrayInput([HttpTrigger(AuthorizationLevel.Anonymous, "POST", Route = "/user/createWithArray")]HttpRequest req, ExecutionContext context)
        {
            var method = this.GetType().GetMethod("CreateUsersWithArrayInput");

            return method != null 
                ? (await ((Task<IActionResult>)method.Invoke(this, new object[] { req, context })).ConfigureAwait(false))
                : new StatusCodeResult((int)HttpStatusCode.NotImplemented);
        }

        [FunctionName("UserApi_CreateUsersWithListInput")]
        public async Task<IActionResult> _CreateUsersWithListInput([HttpTrigger(AuthorizationLevel.Anonymous, "POST", Route = "/user/createWithList")]HttpRequest req, ExecutionContext context)
        {
            var method = this.GetType().GetMethod("CreateUsersWithListInput");

            return method != null 
                ? (await ((Task<IActionResult>)method.Invoke(this, new object[] { req, context })).ConfigureAwait(false))
                : new StatusCodeResult((int)HttpStatusCode.NotImplemented);
        }

        [FunctionName("UserApi_DeleteUser")]
        public async Task<IActionResult> _DeleteUser([HttpTrigger(AuthorizationLevel.Anonymous, "DELETE", Route = "/user/{username}")]HttpRequest req, ExecutionContext context, string username)
        {
            var method = this.GetType().GetMethod("DeleteUser");

            return method != null 
                ? (await ((Task<IActionResult>)method.Invoke(this, new object[] { req, context,  })).ConfigureAwait(false))
                : new StatusCodeResult((int)HttpStatusCode.NotImplemented);
        }

        [FunctionName("UserApi_GetUserByName")]
        public async Task<IActionResult> _GetUserByName([HttpTrigger(AuthorizationLevel.Anonymous, "GET", Route = "/user/{username}")]HttpRequest req, ExecutionContext context, string username)
        {
            var method = this.GetType().GetMethod("GetUserByName");

            return method != null 
                ? (await ((Task<IActionResult>)method.Invoke(this, new object[] { req, context,  })).ConfigureAwait(false))
                : new StatusCodeResult((int)HttpStatusCode.NotImplemented);
        }

        [FunctionName("UserApi_LoginUser")]
        public async Task<IActionResult> _LoginUser([HttpTrigger(AuthorizationLevel.Anonymous, "GET", Route = "/user/login")]HttpRequest req, ExecutionContext context)
        {
            var method = this.GetType().GetMethod("LoginUser");

            return method != null 
                ? (await ((Task<IActionResult>)method.Invoke(this, new object[] { req, context })).ConfigureAwait(false))
                : new StatusCodeResult((int)HttpStatusCode.NotImplemented);
        }

        [FunctionName("UserApi_LogoutUser")]
        public async Task<IActionResult> _LogoutUser([HttpTrigger(AuthorizationLevel.Anonymous, "GET", Route = "/user/logout")]HttpRequest req, ExecutionContext context)
        {
            var method = this.GetType().GetMethod("LogoutUser");

            return method != null 
                ? (await ((Task<IActionResult>)method.Invoke(this, new object[] { req, context })).ConfigureAwait(false))
                : new StatusCodeResult((int)HttpStatusCode.NotImplemented);
        }

        [FunctionName("UserApi_UpdateUser")]
        public async Task<IActionResult> _UpdateUser([HttpTrigger(AuthorizationLevel.Anonymous, "PUT", Route = "/user/{username}")]HttpRequest req, ExecutionContext context, string username)
        {
            var method = this.GetType().GetMethod("UpdateUser");

            return method != null 
                ? (await ((Task<IActionResult>)method.Invoke(this, new object[] { req, context,  })).ConfigureAwait(false))
                : new StatusCodeResult((int)HttpStatusCode.NotImplemented);
        }
    }
}

