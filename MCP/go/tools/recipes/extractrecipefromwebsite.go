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

func ExtractrecipefromwebsiteHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["url"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("url=%v", val))
		}
		if val, ok := args["forceExtraction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("forceExtraction=%v", val))
		}
		if val, ok := args["analyze"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("analyze=%v", val))
		}
		if val, ok := args["includeNutrition"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeNutrition=%v", val))
		}
		if val, ok := args["includeTaste"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeTaste=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/recipes/extract%s", cfg.BaseURL, queryString)
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

func CreateExtractrecipefromwebsiteTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_recipes_extract",
		mcp.WithDescription("Extract Recipe from Website"),
		mcp.WithString("url", mcp.Required(), mcp.Description("The URL of the recipe page.")),
		mcp.WithBoolean("forceExtraction", mcp.Description("If true, the extraction will be triggered whether we already know the recipe or not. Use this only if information is missing as this operation is slower.")),
		mcp.WithBoolean("analyze", mcp.Description("If true, the recipe will be analyzed and classified resolving in more data such as cuisines, dish types, and more.")),
		mcp.WithBoolean("includeNutrition", mcp.Description("Include nutrition data in the recipe information. Nutrition data is per serving. If you want the nutrition data for the entire recipe, just multiply by the number of servings.")),
		mcp.WithBoolean("includeTaste", mcp.Description("Whether taste data should be added to correctly parsed ingredients.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ExtractrecipefromwebsiteHandler(cfg),
	}
}
