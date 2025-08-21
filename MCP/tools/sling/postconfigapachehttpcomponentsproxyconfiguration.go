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

func PostconfigapachehttpcomponentsproxyconfigurationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["proxy.host"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy.host=%v", val))
		}
		if val, ok := args["proxy.host@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy.host@TypeHint=%v", val))
		}
		if val, ok := args["proxy.port"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy.port=%v", val))
		}
		if val, ok := args["proxy.port@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy.port@TypeHint=%v", val))
		}
		if val, ok := args["proxy.exceptions"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy.exceptions=%v", val))
		}
		if val, ok := args["proxy.exceptions@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy.exceptions@TypeHint=%v", val))
		}
		if val, ok := args["proxy.enabled"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy.enabled=%v", val))
		}
		if val, ok := args["proxy.enabled@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy.enabled@TypeHint=%v", val))
		}
		if val, ok := args["proxy.user"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy.user=%v", val))
		}
		if val, ok := args["proxy.user@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy.user@TypeHint=%v", val))
		}
		if val, ok := args["proxy.password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy.password=%v", val))
		}
		if val, ok := args["proxy.password@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("proxy.password@TypeHint=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/apps/system/config/org.apache.http.proxyconfigurator.config%s", cfg.BaseURL, queryString)
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

func CreatePostconfigapachehttpcomponentsproxyconfigurationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_apps_system_config_org.apache.http.proxyconfigurator.config",
		mcp.WithDescription(""),
		mcp.WithString("proxy.host", mcp.Description("")),
		mcp.WithString("proxy.host@TypeHint", mcp.Description("")),
		mcp.WithNumber("proxy.port", mcp.Description("")),
		mcp.WithString("proxy.port@TypeHint", mcp.Description("")),
		mcp.WithArray("proxy.exceptions", mcp.Description("")),
		mcp.WithString("proxy.exceptions@TypeHint", mcp.Description("")),
		mcp.WithBoolean("proxy.enabled", mcp.Description("")),
		mcp.WithString("proxy.enabled@TypeHint", mcp.Description("")),
		mcp.WithString("proxy.user", mcp.Description("")),
		mcp.WithString("proxy.user@TypeHint", mcp.Description("")),
		mcp.WithString("proxy.password", mcp.Description("")),
		mcp.WithString("proxy.password@TypeHint", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PostconfigapachehttpcomponentsproxyconfigurationHandler(cfg),
	}
}
