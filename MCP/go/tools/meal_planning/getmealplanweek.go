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

func GetmealplanweekHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		usernameVal, ok := args["username"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: username"), nil
		}
		username, ok := usernameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: username"), nil
		}
		start_dateVal, ok := args["start-date"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: start-date"), nil
		}
		start_date, ok := start_dateVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: start-date"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["hash"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("hash=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/mealplanner/%s/week/%s%s", cfg.BaseURL, username, start_date, queryString)
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

func CreateGetmealplanweekTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_mealplanner_username_week_start-date",
		mcp.WithDescription("Get Meal Plan Week"),
		mcp.WithString("username", mcp.Required(), mcp.Description("The username.")),
		mcp.WithString("start-date", mcp.Required(), mcp.Description("The start date of the meal planned week in the format yyyy-mm-dd.")),
		mcp.WithString("hash", mcp.Required(), mcp.Description("The private hash for the username.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetmealplanweekHandler(cfg),
	}
}
