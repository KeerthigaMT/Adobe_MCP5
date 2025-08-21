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

func PostagentHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		runmodeVal, ok := args["runmode"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: runmode"), nil
		}
		runmode, ok := runmodeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: runmode"), nil
		}
		nameVal, ok := args["name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: name"), nil
		}
		name, ok := nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: name"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["jcr:content/cq:distribute"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/cq:distribute=%v", val))
		}
		if val, ok := args["jcr:content/cq:distribute@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/cq:distribute@TypeHint=%v", val))
		}
		if val, ok := args["jcr:content/cq:name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/cq:name=%v", val))
		}
		if val, ok := args["jcr:content/cq:template"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/cq:template=%v", val))
		}
		if val, ok := args["jcr:content/enabled"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/enabled=%v", val))
		}
		if val, ok := args["jcr:content/jcr:description"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/jcr:description=%v", val))
		}
		if val, ok := args["jcr:content/jcr:lastModified"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/jcr:lastModified=%v", val))
		}
		if val, ok := args["jcr:content/jcr:lastModifiedBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/jcr:lastModifiedBy=%v", val))
		}
		if val, ok := args["jcr:content/jcr:mixinTypes"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/jcr:mixinTypes=%v", val))
		}
		if val, ok := args["jcr:content/jcr:title"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/jcr:title=%v", val))
		}
		if val, ok := args["jcr:content/logLevel"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/logLevel=%v", val))
		}
		if val, ok := args["jcr:content/noStatusUpdate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/noStatusUpdate=%v", val))
		}
		if val, ok := args["jcr:content/noVersioning"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/noVersioning=%v", val))
		}
		if val, ok := args["jcr:content/protocolConnectTimeout"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/protocolConnectTimeout=%v", val))
		}
		if val, ok := args["jcr:content/protocolHTTPConnectionClosed"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/protocolHTTPConnectionClosed=%v", val))
		}
		if val, ok := args["jcr:content/protocolHTTPExpired"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/protocolHTTPExpired=%v", val))
		}
		if val, ok := args["jcr:content/protocolHTTPHeaders"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/protocolHTTPHeaders=%v", val))
		}
		if val, ok := args["jcr:content/protocolHTTPHeaders@TypeHint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/protocolHTTPHeaders@TypeHint=%v", val))
		}
		if val, ok := args["jcr:content/protocolHTTPMethod"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/protocolHTTPMethod=%v", val))
		}
		if val, ok := args["jcr:content/protocolHTTPSRelaxed"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/protocolHTTPSRelaxed=%v", val))
		}
		if val, ok := args["jcr:content/protocolInterface"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/protocolInterface=%v", val))
		}
		if val, ok := args["jcr:content/protocolSocketTimeout"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/protocolSocketTimeout=%v", val))
		}
		if val, ok := args["jcr:content/protocolVersion"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/protocolVersion=%v", val))
		}
		if val, ok := args["jcr:content/proxyNTLMDomain"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/proxyNTLMDomain=%v", val))
		}
		if val, ok := args["jcr:content/proxyNTLMHost"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/proxyNTLMHost=%v", val))
		}
		if val, ok := args["jcr:content/proxyHost"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/proxyHost=%v", val))
		}
		if val, ok := args["jcr:content/proxyPassword"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/proxyPassword=%v", val))
		}
		if val, ok := args["jcr:content/proxyPort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/proxyPort=%v", val))
		}
		if val, ok := args["jcr:content/proxyUser"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/proxyUser=%v", val))
		}
		if val, ok := args["jcr:content/queueBatchMaxSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/queueBatchMaxSize=%v", val))
		}
		if val, ok := args["jcr:content/queueBatchMode"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/queueBatchMode=%v", val))
		}
		if val, ok := args["jcr:content/queueBatchWaitTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/queueBatchWaitTime=%v", val))
		}
		if val, ok := args["jcr:content/retryDelay"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/retryDelay=%v", val))
		}
		if val, ok := args["jcr:content/reverseReplication"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/reverseReplication=%v", val))
		}
		if val, ok := args["jcr:content/serializationType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/serializationType=%v", val))
		}
		if val, ok := args["jcr:content/sling:resourceType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/sling:resourceType=%v", val))
		}
		if val, ok := args["jcr:content/ssl"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/ssl=%v", val))
		}
		if val, ok := args["jcr:content/transportNTLMDomain"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/transportNTLMDomain=%v", val))
		}
		if val, ok := args["jcr:content/transportNTLMHost"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/transportNTLMHost=%v", val))
		}
		if val, ok := args["jcr:content/transportPassword"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/transportPassword=%v", val))
		}
		if val, ok := args["jcr:content/transportUri"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/transportUri=%v", val))
		}
		if val, ok := args["jcr:content/transportUser"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/transportUser=%v", val))
		}
		if val, ok := args["jcr:content/triggerDistribute"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/triggerDistribute=%v", val))
		}
		if val, ok := args["jcr:content/triggerModified"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/triggerModified=%v", val))
		}
		if val, ok := args["jcr:content/triggerOnOffTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/triggerOnOffTime=%v", val))
		}
		if val, ok := args["jcr:content/triggerReceive"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/triggerReceive=%v", val))
		}
		if val, ok := args["jcr:content/triggerSpecific"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/triggerSpecific=%v", val))
		}
		if val, ok := args["jcr:content/userId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:content/userId=%v", val))
		}
		if val, ok := args["jcr:primaryType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jcr:primaryType=%v", val))
		}
		if val, ok := args[":operation"]; ok {
			queryParams = append(queryParams, fmt.Sprintf(":operation=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/etc/replication/agents.%s/%s%s", cfg.BaseURL, runmode, name, queryString)
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

func CreatePostagentTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_etc_replication_agents.runmode_name",
		mcp.WithDescription(""),
		mcp.WithString("runmode", mcp.Required(), mcp.Description("")),
		mcp.WithString("name", mcp.Required(), mcp.Description("")),
		mcp.WithBoolean("jcr:content/cq:distribute", mcp.Description("")),
		mcp.WithString("jcr:content/cq:distribute@TypeHint", mcp.Description("")),
		mcp.WithString("jcr:content/cq:name", mcp.Description("")),
		mcp.WithString("jcr:content/cq:template", mcp.Description("")),
		mcp.WithBoolean("jcr:content/enabled", mcp.Description("")),
		mcp.WithString("jcr:content/jcr:description", mcp.Description("")),
		mcp.WithString("jcr:content/jcr:lastModified", mcp.Description("")),
		mcp.WithString("jcr:content/jcr:lastModifiedBy", mcp.Description("")),
		mcp.WithString("jcr:content/jcr:mixinTypes", mcp.Description("")),
		mcp.WithString("jcr:content/jcr:title", mcp.Description("")),
		mcp.WithString("jcr:content/logLevel", mcp.Description("")),
		mcp.WithBoolean("jcr:content/noStatusUpdate", mcp.Description("")),
		mcp.WithBoolean("jcr:content/noVersioning", mcp.Description("")),
		mcp.WithString("jcr:content/protocolConnectTimeout", mcp.Description("")),
		mcp.WithBoolean("jcr:content/protocolHTTPConnectionClosed", mcp.Description("")),
		mcp.WithString("jcr:content/protocolHTTPExpired", mcp.Description("")),
		mcp.WithArray("jcr:content/protocolHTTPHeaders", mcp.Description("")),
		mcp.WithString("jcr:content/protocolHTTPHeaders@TypeHint", mcp.Description("")),
		mcp.WithString("jcr:content/protocolHTTPMethod", mcp.Description("")),
		mcp.WithBoolean("jcr:content/protocolHTTPSRelaxed", mcp.Description("")),
		mcp.WithString("jcr:content/protocolInterface", mcp.Description("")),
		mcp.WithString("jcr:content/protocolSocketTimeout", mcp.Description("")),
		mcp.WithString("jcr:content/protocolVersion", mcp.Description("")),
		mcp.WithString("jcr:content/proxyNTLMDomain", mcp.Description("")),
		mcp.WithString("jcr:content/proxyNTLMHost", mcp.Description("")),
		mcp.WithString("jcr:content/proxyHost", mcp.Description("")),
		mcp.WithString("jcr:content/proxyPassword", mcp.Description("")),
		mcp.WithString("jcr:content/proxyPort", mcp.Description("")),
		mcp.WithString("jcr:content/proxyUser", mcp.Description("")),
		mcp.WithString("jcr:content/queueBatchMaxSize", mcp.Description("")),
		mcp.WithString("jcr:content/queueBatchMode", mcp.Description("")),
		mcp.WithString("jcr:content/queueBatchWaitTime", mcp.Description("")),
		mcp.WithString("jcr:content/retryDelay", mcp.Description("")),
		mcp.WithBoolean("jcr:content/reverseReplication", mcp.Description("")),
		mcp.WithString("jcr:content/serializationType", mcp.Description("")),
		mcp.WithString("jcr:content/sling:resourceType", mcp.Description("")),
		mcp.WithString("jcr:content/ssl", mcp.Description("")),
		mcp.WithString("jcr:content/transportNTLMDomain", mcp.Description("")),
		mcp.WithString("jcr:content/transportNTLMHost", mcp.Description("")),
		mcp.WithString("jcr:content/transportPassword", mcp.Description("")),
		mcp.WithString("jcr:content/transportUri", mcp.Description("")),
		mcp.WithString("jcr:content/transportUser", mcp.Description("")),
		mcp.WithBoolean("jcr:content/triggerDistribute", mcp.Description("")),
		mcp.WithBoolean("jcr:content/triggerModified", mcp.Description("")),
		mcp.WithBoolean("jcr:content/triggerOnOffTime", mcp.Description("")),
		mcp.WithBoolean("jcr:content/triggerReceive", mcp.Description("")),
		mcp.WithBoolean("jcr:content/triggerSpecific", mcp.Description("")),
		mcp.WithString("jcr:content/userId", mcp.Description("")),
		mcp.WithString("jcr:primaryType", mcp.Description("")),
		mcp.WithString(":operation", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PostagentHandler(cfg),
	}
}
