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

func PostconfigapacheslinggetservletHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["json.maximumresults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("json.maximumresults=%v", val))
		}
		if val, ok := args["json.maximumresults@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("json.maximumresults@TypeHint=%v", val))
		}
		if val, ok := args["enable.html"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("enable.html=%v", val))
		}
		if val, ok := args["enable.html@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("enable.html@TypeHint=%v", val))
		}
		if val, ok := args["enable.txt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("enable.txt=%v", val))
		}
		if val, ok := args["enable.txt@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("enable.txt@TypeHint=%v", val))
		}
		if val, ok := args["enable.xml"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("enable.xml=%v", val))
		}
		if val, ok := args["enable.xml@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("enable.xml@TypeHint=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/apps/system/config/org.apache.sling.servlets.get.DefaultGetServlet%s", cfg.BaseURL, queryString)
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

func CreatePostconfigapacheslinggetservletTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_apps_system_config_org.apache.sling.servlets.get.DefaultGetServlet",
		mcp.WithDescription(""),
		mcp.WithString("json.maximumresults", mcp.Description("")),
		mcp.WithString("json.maximumresults@TypeHint", mcp.Description("")),
		mcp.WithBoolean("enable.html", mcp.Description("")),
		mcp.WithString("enable.html@TypeHint", mcp.Description("")),
		mcp.WithBoolean("enable.txt", mcp.Description("")),
		mcp.WithString("enable.txt@TypeHint", mcp.Description("")),
		mcp.WithBoolean("enable.xml", mcp.Description("")),
		mcp.WithString("enable.xml@TypeHint", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PostconfigapacheslinggetservletHandler(cfg),
	}
}
