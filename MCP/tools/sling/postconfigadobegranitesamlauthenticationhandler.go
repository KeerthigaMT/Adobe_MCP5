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

func PostconfigadobegranitesamlauthenticationhandlerHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["keyStorePassword"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("keyStorePassword=%v", val))
		}
		if val, ok := args["keyStorePassword@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("keyStorePassword@TypeHint=%v", val))
		}
		if val, ok := args["service.ranking"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("service.ranking=%v", val))
		}
		if val, ok := args["service.ranking@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("service.ranking@TypeHint=%v", val))
		}
		if val, ok := args["idpHttpRedirect"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("idpHttpRedirect=%v", val))
		}
		if val, ok := args["idpHttpRedirect@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("idpHttpRedirect@TypeHint=%v", val))
		}
		if val, ok := args["createUser"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("createUser=%v", val))
		}
		if val, ok := args["createUser@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("createUser@TypeHint=%v", val))
		}
		if val, ok := args["defaultRedirectUrl"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("defaultRedirectUrl=%v", val))
		}
		if val, ok := args["defaultRedirectUrl@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("defaultRedirectUrl@TypeHint=%v", val))
		}
		if val, ok := args["userIDAttribute"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("userIDAttribute=%v", val))
		}
		if val, ok := args["userIDAttribute@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("userIDAttribute@TypeHint=%v", val))
		}
		if val, ok := args["defaultGroups"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("defaultGroups=%v", val))
		}
		if val, ok := args["defaultGroups@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("defaultGroups@TypeHint=%v", val))
		}
		if val, ok := args["idpCertAlias"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("idpCertAlias=%v", val))
		}
		if val, ok := args["idpCertAlias@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("idpCertAlias@TypeHint=%v", val))
		}
		if val, ok := args["addGroupMemberships"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("addGroupMemberships=%v", val))
		}
		if val, ok := args["addGroupMemberships@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("addGroupMemberships@TypeHint=%v", val))
		}
		if val, ok := args["path"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("path=%v", val))
		}
		if val, ok := args["path@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("path@TypeHint=%v", val))
		}
		if val, ok := args["synchronizeAttributes"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("synchronizeAttributes=%v", val))
		}
		if val, ok := args["synchronizeAttributes@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("synchronizeAttributes@TypeHint=%v", val))
		}
		if val, ok := args["clockTolerance"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("clockTolerance=%v", val))
		}
		if val, ok := args["clockTolerance@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("clockTolerance@TypeHint=%v", val))
		}
		if val, ok := args["groupMembershipAttribute"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("groupMembershipAttribute=%v", val))
		}
		if val, ok := args["groupMembershipAttribute@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("groupMembershipAttribute@TypeHint=%v", val))
		}
		if val, ok := args["idpUrl"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("idpUrl=%v", val))
		}
		if val, ok := args["idpUrl@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("idpUrl@TypeHint=%v", val))
		}
		if val, ok := args["logoutUrl"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("logoutUrl=%v", val))
		}
		if val, ok := args["logoutUrl@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("logoutUrl@TypeHint=%v", val))
		}
		if val, ok := args["serviceProviderEntityId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("serviceProviderEntityId=%v", val))
		}
		if val, ok := args["serviceProviderEntityId@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("serviceProviderEntityId@TypeHint=%v", val))
		}
		if val, ok := args["assertionConsumerServiceURL"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("assertionConsumerServiceURL=%v", val))
		}
		if val, ok := args["assertionConsumerServiceURL@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("assertionConsumerServiceURL@TypeHint=%v", val))
		}
		if val, ok := args["handleLogout"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("handleLogout=%v", val))
		}
		if val, ok := args["handleLogout@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("handleLogout@TypeHint=%v", val))
		}
		if val, ok := args["spPrivateKeyAlias"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("spPrivateKeyAlias=%v", val))
		}
		if val, ok := args["spPrivateKeyAlias@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("spPrivateKeyAlias@TypeHint=%v", val))
		}
		if val, ok := args["useEncryption"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("useEncryption=%v", val))
		}
		if val, ok := args["useEncryption@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("useEncryption@TypeHint=%v", val))
		}
		if val, ok := args["nameIdFormat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nameIdFormat=%v", val))
		}
		if val, ok := args["nameIdFormat@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nameIdFormat@TypeHint=%v", val))
		}
		if val, ok := args["digestMethod"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("digestMethod=%v", val))
		}
		if val, ok := args["digestMethod@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("digestMethod@TypeHint=%v", val))
		}
		if val, ok := args["signatureMethod"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("signatureMethod=%v", val))
		}
		if val, ok := args["signatureMethod@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("signatureMethod@TypeHint=%v", val))
		}
		if val, ok := args["userIntermediatePath"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("userIntermediatePath=%v", val))
		}
		if val, ok := args["userIntermediatePath@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("userIntermediatePath@TypeHint=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/apps/system/config/com.adobe.granite.auth.saml.SamlAuthenticationHandler.config%s", cfg.BaseURL, queryString)
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
		var result map[string]interface{}
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

func CreatePostconfigadobegranitesamlauthenticationhandlerTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_apps_system_config_com.adobe.granite.auth.saml.SamlAuthenticationHandler.config",
		mcp.WithDescription(""),
		mcp.WithString("keyStorePassword", mcp.Description("")),
		mcp.WithString("keyStorePassword@TypeHint", mcp.Description("")),
		mcp.WithNumber("service.ranking", mcp.Description("")),
		mcp.WithString("service.ranking@TypeHint", mcp.Description("")),
		mcp.WithBoolean("idpHttpRedirect", mcp.Description("")),
		mcp.WithString("idpHttpRedirect@TypeHint", mcp.Description("")),
		mcp.WithBoolean("createUser", mcp.Description("")),
		mcp.WithString("createUser@TypeHint", mcp.Description("")),
		mcp.WithString("defaultRedirectUrl", mcp.Description("")),
		mcp.WithString("defaultRedirectUrl@TypeHint", mcp.Description("")),
		mcp.WithString("userIDAttribute", mcp.Description("")),
		mcp.WithString("userIDAttribute@TypeHint", mcp.Description("")),
		mcp.WithArray("defaultGroups", mcp.Description("")),
		mcp.WithString("defaultGroups@TypeHint", mcp.Description("")),
		mcp.WithString("idpCertAlias", mcp.Description("")),
		mcp.WithString("idpCertAlias@TypeHint", mcp.Description("")),
		mcp.WithBoolean("addGroupMemberships", mcp.Description("")),
		mcp.WithString("addGroupMemberships@TypeHint", mcp.Description("")),
		mcp.WithArray("path", mcp.Description("")),
		mcp.WithString("path@TypeHint", mcp.Description("")),
		mcp.WithArray("synchronizeAttributes", mcp.Description("")),
		mcp.WithString("synchronizeAttributes@TypeHint", mcp.Description("")),
		mcp.WithNumber("clockTolerance", mcp.Description("")),
		mcp.WithString("clockTolerance@TypeHint", mcp.Description("")),
		mcp.WithString("groupMembershipAttribute", mcp.Description("")),
		mcp.WithString("groupMembershipAttribute@TypeHint", mcp.Description("")),
		mcp.WithString("idpUrl", mcp.Description("")),
		mcp.WithString("idpUrl@TypeHint", mcp.Description("")),
		mcp.WithString("logoutUrl", mcp.Description("")),
		mcp.WithString("logoutUrl@TypeHint", mcp.Description("")),
		mcp.WithString("serviceProviderEntityId", mcp.Description("")),
		mcp.WithString("serviceProviderEntityId@TypeHint", mcp.Description("")),
		mcp.WithString("assertionConsumerServiceURL", mcp.Description("")),
		mcp.WithString("assertionConsumerServiceURL@TypeHint", mcp.Description("")),
		mcp.WithBoolean("handleLogout", mcp.Description("")),
		mcp.WithString("handleLogout@TypeHint", mcp.Description("")),
		mcp.WithString("spPrivateKeyAlias", mcp.Description("")),
		mcp.WithString("spPrivateKeyAlias@TypeHint", mcp.Description("")),
		mcp.WithBoolean("useEncryption", mcp.Description("")),
		mcp.WithString("useEncryption@TypeHint", mcp.Description("")),
		mcp.WithString("nameIdFormat", mcp.Description("")),
		mcp.WithString("nameIdFormat@TypeHint", mcp.Description("")),
		mcp.WithString("digestMethod", mcp.Description("")),
		mcp.WithString("digestMethod@TypeHint", mcp.Description("")),
		mcp.WithString("signatureMethod", mcp.Description("")),
		mcp.WithString("signatureMethod@TypeHint", mcp.Description("")),
		mcp.WithString("userIntermediatePath", mcp.Description("")),
		mcp.WithString("userIntermediatePath@TypeHint", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PostconfigadobegranitesamlauthenticationhandlerHandler(cfg),
	}
}
