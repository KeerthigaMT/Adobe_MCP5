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

func PostconfigapacheslingreferrerfilterHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["allow.empty"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("allow.empty=%v", val))
		}
		if val, ok := args["allow.empty@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("allow.empty@TypeHint=%v", val))
		}
		if val, ok := args["allow.hosts"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("allow.hosts=%v", val))
		}
		if val, ok := args["allow.hosts@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("allow.hosts@TypeHint=%v", val))
		}
		if val, ok := args["allow.hosts.regexp"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("allow.hosts.regexp=%v", val))
		}
		if val, ok := args["allow.hosts.regexp@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("allow.hosts.regexp@TypeHint=%v", val))
		}
		if val, ok := args["filter.methods"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter.methods=%v", val))
		}
		if val, ok := args["filter.methods@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter.methods@TypeHint=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/apps/system/config/org.apache.sling.security.impl.ReferrerFilter%s", cfg.BaseURL, queryString)
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

func CreatePostconfigapacheslingreferrerfilterTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_apps_system_config_org.apache.sling.security.impl.ReferrerFilter",
		mcp.WithDescription(""),
		mcp.WithBoolean("allow.empty", mcp.Description("")),
		mcp.WithString("allow.empty@TypeHint", mcp.Description("")),
		mcp.WithString("allow.hosts", mcp.Description("")),
		mcp.WithString("allow.hosts@TypeHint", mcp.Description("")),
		mcp.WithString("allow.hosts.regexp", mcp.Description("")),
		mcp.WithString("allow.hosts.regexp@TypeHint", mcp.Description("")),
		mcp.WithString("filter.methods", mcp.Description("")),
		mcp.WithString("filter.methods@TypeHint", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PostconfigapacheslingreferrerfilterHandler(cfg),
	}
}
