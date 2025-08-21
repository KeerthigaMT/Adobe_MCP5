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

func GetauthorizablekeystoreHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		intermediatePathVal, ok := args["intermediatePath"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: intermediatePath"), nil
		}
		intermediatePath, ok := intermediatePathVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: intermediatePath"), nil
		}
		authorizableIdVal, ok := args["authorizableId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: authorizableId"), nil
		}
		authorizableId, ok := authorizableIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: authorizableId"), nil
		}
		url := fmt.Sprintf("%s/%s/%s.ks.json", cfg.BaseURL, intermediatePath, authorizableId)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.KeystoreInfo
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

func CreateGetauthorizablekeystoreTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_intermediatePath_authorizableId.ks.json",
		mcp.WithDescription(""),
		mcp.WithString("intermediatePath", mcp.Required(), mcp.Description("")),
		mcp.WithString("authorizableId", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetauthorizablekeystoreHandler(cfg),
	}
}
