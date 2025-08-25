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

func GeneratemealplanHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["timeFrame"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timeFrame=%v", val))
		}
		if val, ok := args["targetCalories"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("targetCalories=%v", val))
		}
		if val, ok := args["diet"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("diet=%v", val))
		}
		if val, ok := args["exclude"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exclude=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/mealplanner/generate%s", cfg.BaseURL, queryString)
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

func CreateGeneratemealplanTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_mealplanner_generate",
		mcp.WithDescription("Generate Meal Plan"),
		mcp.WithString("timeFrame", mcp.Description("Either for one \"day\" or an entire \"week\".")),
		mcp.WithString("targetCalories", mcp.Description("What is the caloric target for one day? The meal plan generator will try to get as close as possible to that goal.")),
		mcp.WithString("diet", mcp.Description("Enter a diet that the meal plan has to adhere to. See a full list of supported diets.")),
		mcp.WithString("exclude", mcp.Description("A comma-separated list of allergens or ingredients that must be excluded.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GeneratemealplanHandler(cfg),
	}
}
