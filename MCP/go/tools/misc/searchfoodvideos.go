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

func SearchfoodvideosHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("query=%v", val))
		}
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["cuisine"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cuisine=%v", val))
		}
		if val, ok := args["diet"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("diet=%v", val))
		}
		if val, ok := args["includeIngredients"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeIngredients=%v", val))
		}
		if val, ok := args["excludeIngredients"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("excludeIngredients=%v", val))
		}
		if val, ok := args["minLength"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minLength=%v", val))
		}
		if val, ok := args["maxLength"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxLength=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["number"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("number=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/food/videos/search%s", cfg.BaseURL, queryString)
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

func CreateSearchfoodvideosTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_food_videos_search",
		mcp.WithDescription("Search Food Videos"),
		mcp.WithString("query", mcp.Description("The (natural language) search query.")),
		mcp.WithString("type", mcp.Description("The type of the recipes. See a full list of supported meal types.")),
		mcp.WithString("cuisine", mcp.Description("The cuisine(s) of the recipes. One or more, comma separated. See a full list of supported cuisines.")),
		mcp.WithString("diet", mcp.Description("The diet for which the recipes must be suitable. See a full list of supported diets.")),
		mcp.WithString("includeIngredients", mcp.Description("A comma-separated list of ingredients that the recipes should contain.")),
		mcp.WithString("excludeIngredients", mcp.Description("A comma-separated list of ingredients or ingredient types that the recipes must not contain.")),
		mcp.WithString("minLength", mcp.Description("Minimum video length in seconds.")),
		mcp.WithString("maxLength", mcp.Description("Maximum video length in seconds.")),
		mcp.WithNumber("offset", mcp.Description("The number of results to skip (between 0 and 900).")),
		mcp.WithNumber("number", mcp.Description("The maximum number of items to return (between 1 and 100). Defaults to 10.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SearchfoodvideosHandler(cfg),
	}
}
