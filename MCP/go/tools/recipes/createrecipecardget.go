package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spoonacular-api/mcp-server/config"
	"github.com/spoonacular-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreaterecipecardgetHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["mask"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("mask=%v", val))
		}
		if val, ok := args["backgroundImage"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("backgroundImage=%v", val))
		}
		if val, ok := args["backgroundColor"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("backgroundColor=%v", val))
		}
		if val, ok := args["fontColor"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fontColor=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/recipes/%s/card%s", cfg.BaseURL, id, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("x-api-key", cfg.APIKey)
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

func CreateCreaterecipecardgetTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_recipes_id_card",
		mcp.WithDescription("Create Recipe Card"),
		mcp.WithString("id", mcp.Required(), mcp.Description("The recipe id.")),
		mcp.WithString("mask", mcp.Description("The mask to put over the recipe image (\"ellipseMask\", \"diamondMask\", \"starMask\", \"heartMask\", \"potMask\", \"fishMask\").")),
		mcp.WithString("backgroundImage", mcp.Description("The background image (\"none\",\"background1\", or \"background2\").")),
		mcp.WithString("backgroundColor", mcp.Description("The background color for the recipe card as a hex-string.")),
		mcp.WithString("fontColor", mcp.Description("The font color for the recipe card as a hex-string.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreaterecipecardgetHandler(cfg),
	}
}
