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

func SearchrecipesbyingredientsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["ingredients"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ingredients=%v", val))
		}
		if val, ok := args["number"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("number=%v", val))
		}
		if val, ok := args["limitLicense"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limitLicense=%v", val))
		}
		if val, ok := args["ranking"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ranking=%v", val))
		}
		if val, ok := args["ignorePantry"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ignorePantry=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/recipes/findByIngredients%s", cfg.BaseURL, queryString)
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
		var result []map[string]interface{}
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

func CreateSearchrecipesbyingredientsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_recipes_findByIngredients",
		mcp.WithDescription("Search Recipes by Ingredients"),
		mcp.WithString("ingredients", mcp.Description("A comma-separated list of ingredients that the recipes should contain.")),
		mcp.WithNumber("number", mcp.Description("The maximum number of items to return (between 1 and 100). Defaults to 10.")),
		mcp.WithBoolean("limitLicense", mcp.Description("Whether the recipes should have an open license that allows display with proper attribution.")),
		mcp.WithString("ranking", mcp.Description("Whether to maximize used ingredients (1) or minimize missing ingredients (2) first.")),
		mcp.WithBoolean("ignorePantry", mcp.Description("Whether to ignore typical pantry items, such as water, salt, flour, etc.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SearchrecipesbyingredientsHandler(cfg),
	}
}
