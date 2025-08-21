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

func PostconfigapachefelixjettybasedhttpserviceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["org.apache.felix.https.nio"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.nio=%v", val))
		}
		if val, ok := args["org.apache.felix.https.nio@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.nio@TypeHint=%v", val))
		}
		if val, ok := args["org.apache.felix.https.keystore"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.keystore=%v", val))
		}
		if val, ok := args["org.apache.felix.https.keystore@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.keystore@TypeHint=%v", val))
		}
		if val, ok := args["org.apache.felix.https.keystore.password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.keystore.password=%v", val))
		}
		if val, ok := args["org.apache.felix.https.keystore.password@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.keystore.password@TypeHint=%v", val))
		}
		if val, ok := args["org.apache.felix.https.keystore.key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.keystore.key=%v", val))
		}
		if val, ok := args["org.apache.felix.https.keystore.key@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.keystore.key@TypeHint=%v", val))
		}
		if val, ok := args["org.apache.felix.https.keystore.key.password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.keystore.key.password=%v", val))
		}
		if val, ok := args["org.apache.felix.https.keystore.key.password@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.keystore.key.password@TypeHint=%v", val))
		}
		if val, ok := args["org.apache.felix.https.truststore"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.truststore=%v", val))
		}
		if val, ok := args["org.apache.felix.https.truststore@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.truststore@TypeHint=%v", val))
		}
		if val, ok := args["org.apache.felix.https.truststore.password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.truststore.password=%v", val))
		}
		if val, ok := args["org.apache.felix.https.truststore.password@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.truststore.password@TypeHint=%v", val))
		}
		if val, ok := args["org.apache.felix.https.clientcertificate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.clientcertificate=%v", val))
		}
		if val, ok := args["org.apache.felix.https.clientcertificate@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.clientcertificate@TypeHint=%v", val))
		}
		if val, ok := args["org.apache.felix.https.enable"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.enable=%v", val))
		}
		if val, ok := args["org.apache.felix.https.enable@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.apache.felix.https.enable@TypeHint=%v", val))
		}
		if val, ok := args["org.osgi.service.http.port.secure"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.osgi.service.http.port.secure=%v", val))
		}
		if val, ok := args["org.osgi.service.http.port.secure@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org.osgi.service.http.port.secure@TypeHint=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/apps/system/config/org.apache.felix.http%s", cfg.BaseURL, queryString)
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

func CreatePostconfigapachefelixjettybasedhttpserviceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_apps_system_config_org.apache.felix.http",
		mcp.WithDescription(""),
		mcp.WithBoolean("org.apache.felix.https.nio", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.nio@TypeHint", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.keystore", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.keystore@TypeHint", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.keystore.password", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.keystore.password@TypeHint", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.keystore.key", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.keystore.key@TypeHint", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.keystore.key.password", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.keystore.key.password@TypeHint", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.truststore", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.truststore@TypeHint", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.truststore.password", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.truststore.password@TypeHint", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.clientcertificate", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.clientcertificate@TypeHint", mcp.Description("")),
		mcp.WithBoolean("org.apache.felix.https.enable", mcp.Description("")),
		mcp.WithString("org.apache.felix.https.enable@TypeHint", mcp.Description("")),
		mcp.WithString("org.osgi.service.http.port.secure", mcp.Description("")),
		mcp.WithString("org.osgi.service.http.port.secure@TypeHint", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PostconfigapachefelixjettybasedhttpserviceHandler(cfg),
	}
}
