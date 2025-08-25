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

func IngredientsearchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("query=%v", val))
		}
		if val, ok := args["addChildren"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("addChildren=%v", val))
		}
		if val, ok := args["minProteinPercent"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minProteinPercent=%v", val))
		}
		if val, ok := args["maxProteinPercent"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxProteinPercent=%v", val))
		}
		if val, ok := args["minFatPercent"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minFatPercent=%v", val))
		}
		if val, ok := args["maxFatPercent"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxFatPercent=%v", val))
		}
		if val, ok := args["minCarbsPercent"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minCarbsPercent=%v", val))
		}
		if val, ok := args["maxCarbsPercent"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxCarbsPercent=%v", val))
		}
		if val, ok := args["metaInformation"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("metaInformation=%v", val))
		}
		if val, ok := args["intolerances"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("intolerances=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["sortDirection"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sortDirection=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["number"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("number=%v", val))
		}
		if val, ok := args["language"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("language=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/food/ingredients/search%s", cfg.BaseURL, queryString)
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

func CreateIngredientsearchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_food_ingredients_search",
		mcp.WithDescription("Ingredient Search"),
		mcp.WithString("query", mcp.Description("The (natural language) search query.")),
		mcp.WithBoolean("addChildren", mcp.Description("Whether to add children of found foods.")),
		mcp.WithString("minProteinPercent", mcp.Description("The minimum percentage of protein the food must have (between 0 and 100).")),
		mcp.WithString("maxProteinPercent", mcp.Description("The maximum percentage of protein the food can have (between 0 and 100).")),
		mcp.WithString("minFatPercent", mcp.Description("The minimum percentage of fat the food must have (between 0 and 100).")),
		mcp.WithString("maxFatPercent", mcp.Description("The maximum percentage of fat the food can have (between 0 and 100).")),
		mcp.WithString("minCarbsPercent", mcp.Description("The minimum percentage of carbs the food must have (between 0 and 100).")),
		mcp.WithString("maxCarbsPercent", mcp.Description("The maximum percentage of carbs the food can have (between 0 and 100).")),
		mcp.WithBoolean("metaInformation", mcp.Description("Whether to return more meta information about the ingredients.")),
		mcp.WithString("intolerances", mcp.Description("A comma-separated list of intolerances. All recipes returned must not contain ingredients that are not suitable for people with the intolerances entered. See a full list of supported intolerances.")),
		mcp.WithString("sort", mcp.Description("The strategy to sort recipes by. See a full list of supported sorting options.")),
		mcp.WithString("sortDirection", mcp.Description("The direction in which to sort. Must be either 'asc' (ascending) or 'desc' (descending).")),
		mcp.WithNumber("offset", mcp.Description("The number of results to skip (between 0 and 900).")),
		mcp.WithNumber("number", mcp.Description("The maximum number of items to return (between 1 and 100). Defaults to 10.")),
		mcp.WithString("language", mcp.Description("The language of the input. Either 'en' or 'de'.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    IngredientsearchHandler(cfg),
	}
}
