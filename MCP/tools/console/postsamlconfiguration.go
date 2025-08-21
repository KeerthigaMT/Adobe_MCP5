package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/adobe-experience-manager-aem-api/mcp-server/config"
	"github.com/adobe-experience-manager-aem-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func PostsamlconfigurationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["post"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("post=%v", val))
		}
		if val, ok := args["apply"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apply=%v", val))
		}
		if val, ok := args["delete"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("delete=%v", val))
		}
		if val, ok := args["action"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("action=%v", val))
		}
		if val, ok := args["$location"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("$location=%v", val))
		}
		if val, ok := args["path"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("path=%v", val))
		}
		if val, ok := args["service.ranking"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("service.ranking=%v", val))
		}
		if val, ok := args["idpUrl"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("idpUrl=%v", val))
		}
		if val, ok := args["idpCertAlias"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("idpCertAlias=%v", val))
		}
		if val, ok := args["idpHttpRedirect"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("idpHttpRedirect=%v", val))
		}
		if val, ok := args["serviceProviderEntityId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("serviceProviderEntityId=%v", val))
		}
		if val, ok := args["assertionConsumerServiceURL"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("assertionConsumerServiceURL=%v", val))
		}
		if val, ok := args["spPrivateKeyAlias"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("spPrivateKeyAlias=%v", val))
		}
		if val, ok := args["keyStorePassword"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("keyStorePassword=%v", val))
		}
		if val, ok := args["defaultRedirectUrl"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("defaultRedirectUrl=%v", val))
		}
		if val, ok := args["userIDAttribute"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("userIDAttribute=%v", val))
		}
		if val, ok := args["useEncryption"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("useEncryption=%v", val))
		}
		if val, ok := args["createUser"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("createUser=%v", val))
		}
		if val, ok := args["addGroupMemberships"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("addGroupMemberships=%v", val))
		}
		if val, ok := args["groupMembershipAttribute"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("groupMembershipAttribute=%v", val))
		}
		if val, ok := args["defaultGroups"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("defaultGroups=%v", val))
		}
		if val, ok := args["nameIdFormat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nameIdFormat=%v", val))
		}
		if val, ok := args["synchronizeAttributes"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("synchronizeAttributes=%v", val))
		}
		if val, ok := args["handleLogout"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("handleLogout=%v", val))
		}
		if val, ok := args["logoutUrl"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("logoutUrl=%v", val))
		}
		if val, ok := args["clockTolerance"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("clockTolerance=%v", val))
		}
		if val, ok := args["digestMethod"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("digestMethod=%v", val))
		}
		if val, ok := args["signatureMethod"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("signatureMethod=%v", val))
		}
		if val, ok := args["userIntermediatePath"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("userIntermediatePath=%v", val))
		}
		if val, ok := args["propertylist"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("propertylist=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/system/console/configMgr/com.adobe.granite.auth.saml.SamlAuthenticationHandler%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BasicAuth != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Basic %s", cfg.BasicAuth))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.SamlConfigurationInfo
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreatePostsamlconfigurationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_system_console_configMgr_com.adobe.granite.auth.saml.SamlAuthenticationHandler",
		mcp.WithDescription(""),
		mcp.WithBoolean("post", mcp.Description("")),
		mcp.WithBoolean("apply", mcp.Description("")),
		mcp.WithBoolean("delete", mcp.Description("")),
		mcp.WithString("action", mcp.Description("")),
		mcp.WithString("$location", mcp.Description("")),
		mcp.WithArray("path", mcp.Description("")),
		mcp.WithNumber("service.ranking", mcp.Description("")),
		mcp.WithString("idpUrl", mcp.Description("")),
		mcp.WithString("idpCertAlias", mcp.Description("")),
		mcp.WithBoolean("idpHttpRedirect", mcp.Description("")),
		mcp.WithString("serviceProviderEntityId", mcp.Description("")),
		mcp.WithString("assertionConsumerServiceURL", mcp.Description("")),
		mcp.WithString("spPrivateKeyAlias", mcp.Description("")),
		mcp.WithString("keyStorePassword", mcp.Description("")),
		mcp.WithString("defaultRedirectUrl", mcp.Description("")),
		mcp.WithString("userIDAttribute", mcp.Description("")),
		mcp.WithBoolean("useEncryption", mcp.Description("")),
		mcp.WithBoolean("createUser", mcp.Description("")),
		mcp.WithBoolean("addGroupMemberships", mcp.Description("")),
		mcp.WithString("groupMembershipAttribute", mcp.Description("")),
		mcp.WithArray("defaultGroups", mcp.Description("")),
		mcp.WithString("nameIdFormat", mcp.Description("")),
		mcp.WithArray("synchronizeAttributes", mcp.Description("")),
		mcp.WithBoolean("handleLogout", mcp.Description("")),
		mcp.WithString("logoutUrl", mcp.Description("")),
		mcp.WithNumber("clockTolerance", mcp.Description("")),
		mcp.WithString("digestMethod", mcp.Description("")),
		mcp.WithString("signatureMethod", mcp.Description("")),
		mcp.WithString("userIntermediatePath", mcp.Description("")),
		mcp.WithArray("propertylist", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PostsamlconfigurationHandler(cfg),
	}
}
