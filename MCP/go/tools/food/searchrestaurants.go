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

func SearchrestaurantsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("query=%v", val))
		}
		if val, ok := args["lat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lat=%v", val))
		}
		if val, ok := args["lng"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lng=%v", val))
		}
		if val, ok := args["distance"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("distance=%v", val))
		}
		if val, ok := args["budget"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("budget=%v", val))
		}
		if val, ok := args["cuisine"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cuisine=%v", val))
		}
		if val, ok := args["min-rating"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("min-rating=%v", val))
		}
		if val, ok := args["is-open"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is-open=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/food/restaurants/search%s", cfg.BaseURL, queryString)
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

func CreateSearchrestaurantsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_food_restaurants_search",
		mcp.WithDescription("Search Restaurants"),
		mcp.WithString("query", mcp.Description("The search query.")),
		mcp.WithString("lat", mcp.Description("The latitude of the user's location.")),
		mcp.WithString("lng", mcp.Description("The longitude of the user's location.\".")),
		mcp.WithString("distance", mcp.Description("The distance around the location in miles.")),
		mcp.WithString("budget", mcp.Description("The user's budget for a meal in USD.")),
		mcp.WithString("cuisine", mcp.Description("The cuisine of the restaurant.")),
		mcp.WithString("min-rating", mcp.Description("The minimum rating of the restaurant between 0 and 5.")),
		mcp.WithBoolean("is-open", mcp.Description("Whether the restaurant must be open at the time of search.")),
		mcp.WithString("sort", mcp.Description("How to sort the results, one of the following 'cheapest', 'fastest', 'rating', 'distance' or the default 'relevance'.")),
		mcp.WithString("page", mcp.Description("The page number of results.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SearchrestaurantsHandler(cfg),
	}
}
