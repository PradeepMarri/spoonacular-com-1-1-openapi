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

func SearchgroceryproductsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("query=%v", val))
		}
		if val, ok := args["minCalories"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minCalories=%v", val))
		}
		if val, ok := args["maxCalories"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxCalories=%v", val))
		}
		if val, ok := args["minCarbs"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minCarbs=%v", val))
		}
		if val, ok := args["maxCarbs"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxCarbs=%v", val))
		}
		if val, ok := args["minProtein"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minProtein=%v", val))
		}
		if val, ok := args["maxProtein"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxProtein=%v", val))
		}
		if val, ok := args["minFat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minFat=%v", val))
		}
		if val, ok := args["maxFat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxFat=%v", val))
		}
		if val, ok := args["addProductInformation"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("addProductInformation=%v", val))
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
		url := fmt.Sprintf("%s/food/products/search%s", cfg.BaseURL, queryString)
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

func CreateSearchgroceryproductsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_food_products_search",
		mcp.WithDescription("Search Grocery Products"),
		mcp.WithString("query", mcp.Description("The (natural language) search query.")),
		mcp.WithString("minCalories", mcp.Description("The minimum amount of calories the product must have.")),
		mcp.WithString("maxCalories", mcp.Description("The maximum amount of calories the product can have.")),
		mcp.WithString("minCarbs", mcp.Description("The minimum amount of carbohydrates in grams the product must have.")),
		mcp.WithString("maxCarbs", mcp.Description("The maximum amount of carbohydrates in grams the product can have.")),
		mcp.WithString("minProtein", mcp.Description("The minimum amount of protein in grams the product must have.")),
		mcp.WithString("maxProtein", mcp.Description("The maximum amount of protein in grams the product can have.")),
		mcp.WithString("minFat", mcp.Description("The minimum amount of fat in grams the product must have.")),
		mcp.WithString("maxFat", mcp.Description("The maximum amount of fat in grams the product can have.")),
		mcp.WithBoolean("addProductInformation", mcp.Description("If set to true, you get more information about the products returned.")),
		mcp.WithNumber("offset", mcp.Description("The number of results to skip (between 0 and 900).")),
		mcp.WithNumber("number", mcp.Description("The maximum number of items to return (between 1 and 100). Defaults to 10.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SearchgroceryproductsHandler(cfg),
	}
}
