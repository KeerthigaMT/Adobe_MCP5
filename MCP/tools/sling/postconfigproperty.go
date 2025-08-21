package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/adobe-experience-manager-aem-api/mcp-server/config"
	"github.com/adobe-experience-manager-aem-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func PostconfigpropertyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		configNodeNameVal, ok := args["configNodeName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: configNodeName"), nil
		}
		configNodeName, ok := configNodeNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: configNodeName"), nil
		}
		url := fmt.Sprintf("%s/apps/system/config/%s", cfg.BaseURL, configNodeName)
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

func CreatePostconfigpropertyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_apps_system_config_configNodeName",
		mcp.WithDescription(""),
		mcp.WithString("configNodeName", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PostconfigpropertyHandler(cfg),
	}
}
