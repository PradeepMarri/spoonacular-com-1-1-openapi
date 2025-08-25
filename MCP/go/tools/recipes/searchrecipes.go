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

func SearchrecipesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("query=%v", val))
		}
		if val, ok := args["cuisine"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cuisine=%v", val))
		}
		if val, ok := args["excludeCuisine"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("excludeCuisine=%v", val))
		}
		if val, ok := args["diet"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("diet=%v", val))
		}
		if val, ok := args["intolerances"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("intolerances=%v", val))
		}
		if val, ok := args["equipment"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("equipment=%v", val))
		}
		if val, ok := args["includeIngredients"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeIngredients=%v", val))
		}
		if val, ok := args["excludeIngredients"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("excludeIngredients=%v", val))
		}
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["instructionsRequired"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("instructionsRequired=%v", val))
		}
		if val, ok := args["fillIngredients"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fillIngredients=%v", val))
		}
		if val, ok := args["addRecipeInformation"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("addRecipeInformation=%v", val))
		}
		if val, ok := args["addRecipeNutrition"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("addRecipeNutrition=%v", val))
		}
		if val, ok := args["author"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("author=%v", val))
		}
		if val, ok := args["tags"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tags=%v", val))
		}
		if val, ok := args["recipeBoxId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("recipeBoxId=%v", val))
		}
		if val, ok := args["titleMatch"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("titleMatch=%v", val))
		}
		if val, ok := args["maxReadyTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxReadyTime=%v", val))
		}
		if val, ok := args["ignorePantry"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ignorePantry=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["sortDirection"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sortDirection=%v", val))
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
		if val, ok := args["minCalories"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minCalories=%v", val))
		}
		if val, ok := args["maxCalories"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxCalories=%v", val))
		}
		if val, ok := args["minFat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minFat=%v", val))
		}
		if val, ok := args["maxFat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxFat=%v", val))
		}
		if val, ok := args["minAlcohol"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minAlcohol=%v", val))
		}
		if val, ok := args["maxAlcohol"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxAlcohol=%v", val))
		}
		if val, ok := args["minCaffeine"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minCaffeine=%v", val))
		}
		if val, ok := args["maxCaffeine"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxCaffeine=%v", val))
		}
		if val, ok := args["minCopper"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minCopper=%v", val))
		}
		if val, ok := args["maxCopper"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxCopper=%v", val))
		}
		if val, ok := args["minCalcium"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minCalcium=%v", val))
		}
		if val, ok := args["maxCalcium"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxCalcium=%v", val))
		}
		if val, ok := args["minCholine"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minCholine=%v", val))
		}
		if val, ok := args["maxCholine"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxCholine=%v", val))
		}
		if val, ok := args["minCholesterol"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minCholesterol=%v", val))
		}
		if val, ok := args["maxCholesterol"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxCholesterol=%v", val))
		}
		if val, ok := args["minFluoride"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minFluoride=%v", val))
		}
		if val, ok := args["maxFluoride"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxFluoride=%v", val))
		}
		if val, ok := args["minSaturatedFat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minSaturatedFat=%v", val))
		}
		if val, ok := args["maxSaturatedFat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxSaturatedFat=%v", val))
		}
		if val, ok := args["minVitaminA"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minVitaminA=%v", val))
		}
		if val, ok := args["maxVitaminA"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxVitaminA=%v", val))
		}
		if val, ok := args["minVitaminC"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minVitaminC=%v", val))
		}
		if val, ok := args["maxVitaminC"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxVitaminC=%v", val))
		}
		if val, ok := args["minVitaminD"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minVitaminD=%v", val))
		}
		if val, ok := args["maxVitaminD"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxVitaminD=%v", val))
		}
		if val, ok := args["minVitaminE"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minVitaminE=%v", val))
		}
		if val, ok := args["maxVitaminE"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxVitaminE=%v", val))
		}
		if val, ok := args["minVitaminK"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minVitaminK=%v", val))
		}
		if val, ok := args["maxVitaminK"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxVitaminK=%v", val))
		}
		if val, ok := args["minVitaminB1"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minVitaminB1=%v", val))
		}
		if val, ok := args["maxVitaminB1"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxVitaminB1=%v", val))
		}
		if val, ok := args["minVitaminB2"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minVitaminB2=%v", val))
		}
		if val, ok := args["maxVitaminB2"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxVitaminB2=%v", val))
		}
		if val, ok := args["minVitaminB5"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minVitaminB5=%v", val))
		}
		if val, ok := args["maxVitaminB5"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxVitaminB5=%v", val))
		}
		if val, ok := args["minVitaminB3"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minVitaminB3=%v", val))
		}
		if val, ok := args["maxVitaminB3"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxVitaminB3=%v", val))
		}
		if val, ok := args["minVitaminB6"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minVitaminB6=%v", val))
		}
		if val, ok := args["maxVitaminB6"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxVitaminB6=%v", val))
		}
		if val, ok := args["minVitaminB12"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minVitaminB12=%v", val))
		}
		if val, ok := args["maxVitaminB12"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxVitaminB12=%v", val))
		}
		if val, ok := args["minFiber"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minFiber=%v", val))
		}
		if val, ok := args["maxFiber"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxFiber=%v", val))
		}
		if val, ok := args["minFolate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minFolate=%v", val))
		}
		if val, ok := args["maxFolate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxFolate=%v", val))
		}
		if val, ok := args["minFolicAcid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minFolicAcid=%v", val))
		}
		if val, ok := args["maxFolicAcid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxFolicAcid=%v", val))
		}
		if val, ok := args["minIodine"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minIodine=%v", val))
		}
		if val, ok := args["maxIodine"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxIodine=%v", val))
		}
		if val, ok := args["minIron"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minIron=%v", val))
		}
		if val, ok := args["maxIron"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxIron=%v", val))
		}
		if val, ok := args["minMagnesium"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minMagnesium=%v", val))
		}
		if val, ok := args["maxMagnesium"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxMagnesium=%v", val))
		}
		if val, ok := args["minManganese"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minManganese=%v", val))
		}
		if val, ok := args["maxManganese"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxManganese=%v", val))
		}
		if val, ok := args["minPhosphorus"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minPhosphorus=%v", val))
		}
		if val, ok := args["maxPhosphorus"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxPhosphorus=%v", val))
		}
		if val, ok := args["minPotassium"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minPotassium=%v", val))
		}
		if val, ok := args["maxPotassium"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxPotassium=%v", val))
		}
		if val, ok := args["minSelenium"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minSelenium=%v", val))
		}
		if val, ok := args["maxSelenium"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxSelenium=%v", val))
		}
		if val, ok := args["minSodium"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minSodium=%v", val))
		}
		if val, ok := args["maxSodium"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxSodium=%v", val))
		}
		if val, ok := args["minSugar"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minSugar=%v", val))
		}
		if val, ok := args["maxSugar"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxSugar=%v", val))
		}
		if val, ok := args["minZinc"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minZinc=%v", val))
		}
		if val, ok := args["maxZinc"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxZinc=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["number"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("number=%v", val))
		}
		if val, ok := args["limitLicense"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limitLicense=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/recipes/complexSearch%s", cfg.BaseURL, queryString)
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

func CreateSearchrecipesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_recipes_complexSearch",
		mcp.WithDescription("Search Recipes"),
		mcp.WithString("query", mcp.Description("The (natural language) search query.")),
		mcp.WithString("cuisine", mcp.Description("The cuisine(s) of the recipes. One or more, comma separated (will be interpreted as 'OR'). See a full list of supported cuisines.")),
		mcp.WithString("excludeCuisine", mcp.Description("The cuisine(s) the recipes must not match. One or more, comma separated (will be interpreted as 'AND'). See a full list of supported cuisines.")),
		mcp.WithString("diet", mcp.Description("The diet for which the recipes must be suitable. See a full list of supported diets.")),
		mcp.WithString("intolerances", mcp.Description("A comma-separated list of intolerances. All recipes returned must not contain ingredients that are not suitable for people with the intolerances entered. See a full list of supported intolerances.")),
		mcp.WithString("equipment", mcp.Description("The equipment required. Multiple values will be interpreted as 'or'. For example, value could be \"blender, frying pan, bowl\".")),
		mcp.WithString("includeIngredients", mcp.Description("A comma-separated list of ingredients that should/must be used in the recipes.")),
		mcp.WithString("excludeIngredients", mcp.Description("A comma-separated list of ingredients or ingredient types that the recipes must not contain.")),
		mcp.WithString("type", mcp.Description("The type of recipe. See a full list of supported meal types.")),
		mcp.WithBoolean("instructionsRequired", mcp.Description("Whether the recipes must have instructions.")),
		mcp.WithBoolean("fillIngredients", mcp.Description("Add information about the ingredients and whether they are used or missing in relation to the query.")),
		mcp.WithBoolean("addRecipeInformation", mcp.Description("If set to true, you get more information about the recipes returned.")),
		mcp.WithBoolean("addRecipeNutrition", mcp.Description("If set to true, you get nutritional information about each recipes returned.")),
		mcp.WithString("author", mcp.Description("The username of the recipe author.")),
		mcp.WithString("tags", mcp.Description("The tags (can be diets, meal types, cuisines, or intolerances) that the recipe must have.")),
		mcp.WithString("recipeBoxId", mcp.Description("The id of the recipe box to which the search should be limited to.")),
		mcp.WithString("titleMatch", mcp.Description("Enter text that must be found in the title of the recipes.")),
		mcp.WithString("maxReadyTime", mcp.Description("The maximum time in minutes it should take to prepare and cook the recipe.")),
		mcp.WithBoolean("ignorePantry", mcp.Description("Whether to ignore typical pantry items, such as water, salt, flour, etc.")),
		mcp.WithString("sort", mcp.Description("The strategy to sort recipes by. See a full list of supported sorting options.")),
		mcp.WithString("sortDirection", mcp.Description("The direction in which to sort. Must be either 'asc' (ascending) or 'desc' (descending).")),
		mcp.WithString("minCarbs", mcp.Description("The minimum amount of carbohydrates in grams the recipe must have.")),
		mcp.WithString("maxCarbs", mcp.Description("The maximum amount of carbohydrates in grams the recipe can have.")),
		mcp.WithString("minProtein", mcp.Description("The minimum amount of protein in grams the recipe must have.")),
		mcp.WithString("maxProtein", mcp.Description("The maximum amount of protein in grams the recipe can have.")),
		mcp.WithString("minCalories", mcp.Description("The minimum amount of calories the recipe must have.")),
		mcp.WithString("maxCalories", mcp.Description("The maximum amount of calories the recipe can have.")),
		mcp.WithString("minFat", mcp.Description("The minimum amount of fat in grams the recipe must have.")),
		mcp.WithString("maxFat", mcp.Description("The maximum amount of fat in grams the recipe can have.")),
		mcp.WithString("minAlcohol", mcp.Description("The minimum amount of alcohol in grams the recipe must have.")),
		mcp.WithString("maxAlcohol", mcp.Description("The maximum amount of alcohol in grams the recipe can have.")),
		mcp.WithString("minCaffeine", mcp.Description("The minimum amount of caffeine in milligrams the recipe must have.")),
		mcp.WithString("maxCaffeine", mcp.Description("The maximum amount of caffeine in milligrams the recipe can have.")),
		mcp.WithString("minCopper", mcp.Description("The minimum amount of copper in milligrams the recipe must have.")),
		mcp.WithString("maxCopper", mcp.Description("The maximum amount of copper in milligrams the recipe can have.")),
		mcp.WithString("minCalcium", mcp.Description("The minimum amount of calcium in milligrams the recipe must have.")),
		mcp.WithString("maxCalcium", mcp.Description("The maximum amount of calcium in milligrams the recipe can have.")),
		mcp.WithString("minCholine", mcp.Description("The minimum amount of choline in milligrams the recipe must have.")),
		mcp.WithString("maxCholine", mcp.Description("The maximum amount of choline in milligrams the recipe can have.")),
		mcp.WithString("minCholesterol", mcp.Description("The minimum amount of cholesterol in milligrams the recipe must have.")),
		mcp.WithString("maxCholesterol", mcp.Description("The maximum amount of cholesterol in milligrams the recipe can have.")),
		mcp.WithString("minFluoride", mcp.Description("The minimum amount of fluoride in milligrams the recipe must have.")),
		mcp.WithString("maxFluoride", mcp.Description("The maximum amount of fluoride in milligrams the recipe can have.")),
		mcp.WithString("minSaturatedFat", mcp.Description("The minimum amount of saturated fat in grams the recipe must have.")),
		mcp.WithString("maxSaturatedFat", mcp.Description("The maximum amount of saturated fat in grams the recipe can have.")),
		mcp.WithString("minVitaminA", mcp.Description("The minimum amount of Vitamin A in IU the recipe must have.")),
		mcp.WithString("maxVitaminA", mcp.Description("The maximum amount of Vitamin A in IU the recipe can have.")),
		mcp.WithString("minVitaminC", mcp.Description("The minimum amount of Vitamin C milligrams the recipe must have.")),
		mcp.WithString("maxVitaminC", mcp.Description("The maximum amount of Vitamin C in milligrams the recipe can have.")),
		mcp.WithString("minVitaminD", mcp.Description("The minimum amount of Vitamin D in micrograms the recipe must have.")),
		mcp.WithString("maxVitaminD", mcp.Description("The maximum amount of Vitamin D in micrograms the recipe can have.")),
		mcp.WithString("minVitaminE", mcp.Description("The minimum amount of Vitamin E in milligrams the recipe must have.")),
		mcp.WithString("maxVitaminE", mcp.Description("The maximum amount of Vitamin E in milligrams the recipe can have.")),
		mcp.WithString("minVitaminK", mcp.Description("The minimum amount of Vitamin K in micrograms the recipe must have.")),
		mcp.WithString("maxVitaminK", mcp.Description("The maximum amount of Vitamin K in micrograms the recipe can have.")),
		mcp.WithString("minVitaminB1", mcp.Description("The minimum amount of Vitamin B1 in milligrams the recipe must have.")),
		mcp.WithString("maxVitaminB1", mcp.Description("The maximum amount of Vitamin B1 in milligrams the recipe can have.")),
		mcp.WithString("minVitaminB2", mcp.Description("The minimum amount of Vitamin B2 in milligrams the recipe must have.")),
		mcp.WithString("maxVitaminB2", mcp.Description("The maximum amount of Vitamin B2 in milligrams the recipe can have.")),
		mcp.WithString("minVitaminB5", mcp.Description("The minimum amount of Vitamin B5 in milligrams the recipe must have.")),
		mcp.WithString("maxVitaminB5", mcp.Description("The maximum amount of Vitamin B5 in milligrams the recipe can have.")),
		mcp.WithString("minVitaminB3", mcp.Description("The minimum amount of Vitamin B3 in milligrams the recipe must have.")),
		mcp.WithString("maxVitaminB3", mcp.Description("The maximum amount of Vitamin B3 in milligrams the recipe can have.")),
		mcp.WithString("minVitaminB6", mcp.Description("The minimum amount of Vitamin B6 in milligrams the recipe must have.")),
		mcp.WithString("maxVitaminB6", mcp.Description("The maximum amount of Vitamin B6 in milligrams the recipe can have.")),
		mcp.WithString("minVitaminB12", mcp.Description("The minimum amount of Vitamin B12 in micrograms the recipe must have.")),
		mcp.WithString("maxVitaminB12", mcp.Description("The maximum amount of Vitamin B12 in micrograms the recipe can have.")),
		mcp.WithString("minFiber", mcp.Description("The minimum amount of fiber in grams the recipe must have.")),
		mcp.WithString("maxFiber", mcp.Description("The maximum amount of fiber in grams the recipe can have.")),
		mcp.WithString("minFolate", mcp.Description("The minimum amount of folate in micrograms the recipe must have.")),
		mcp.WithString("maxFolate", mcp.Description("The maximum amount of folate in micrograms the recipe can have.")),
		mcp.WithString("minFolicAcid", mcp.Description("The minimum amount of folic acid in micrograms the recipe must have.")),
		mcp.WithString("maxFolicAcid", mcp.Description("The maximum amount of folic acid in micrograms the recipe can have.")),
		mcp.WithString("minIodine", mcp.Description("The minimum amount of iodine in micrograms the recipe must have.")),
		mcp.WithString("maxIodine", mcp.Description("The maximum amount of iodine in micrograms the recipe can have.")),
		mcp.WithString("minIron", mcp.Description("The minimum amount of iron in milligrams the recipe must have.")),
		mcp.WithString("maxIron", mcp.Description("The maximum amount of iron in milligrams the recipe can have.")),
		mcp.WithString("minMagnesium", mcp.Description("The minimum amount of magnesium in milligrams the recipe must have.")),
		mcp.WithString("maxMagnesium", mcp.Description("The maximum amount of magnesium in milligrams the recipe can have.")),
		mcp.WithString("minManganese", mcp.Description("The minimum amount of manganese in milligrams the recipe must have.")),
		mcp.WithString("maxManganese", mcp.Description("The maximum amount of manganese in milligrams the recipe can have.")),
		mcp.WithString("minPhosphorus", mcp.Description("The minimum amount of phosphorus in milligrams the recipe must have.")),
		mcp.WithString("maxPhosphorus", mcp.Description("The maximum amount of phosphorus in milligrams the recipe can have.")),
		mcp.WithString("minPotassium", mcp.Description("The minimum amount of potassium in milligrams the recipe must have.")),
		mcp.WithString("maxPotassium", mcp.Description("The maximum amount of potassium in milligrams the recipe can have.")),
		mcp.WithString("minSelenium", mcp.Description("The minimum amount of selenium in micrograms the recipe must have.")),
		mcp.WithString("maxSelenium", mcp.Description("The maximum amount of selenium in micrograms the recipe can have.")),
		mcp.WithString("minSodium", mcp.Description("The minimum amount of sodium in milligrams the recipe must have.")),
		mcp.WithString("maxSodium", mcp.Description("The maximum amount of sodium in milligrams the recipe can have.")),
		mcp.WithString("minSugar", mcp.Description("The minimum amount of sugar in grams the recipe must have.")),
		mcp.WithString("maxSugar", mcp.Description("The maximum amount of sugar in grams the recipe can have.")),
		mcp.WithString("minZinc", mcp.Description("The minimum amount of zinc in milligrams the recipe must have.")),
		mcp.WithString("maxZinc", mcp.Description("The maximum amount of zinc in milligrams the recipe can have.")),
		mcp.WithNumber("offset", mcp.Description("The number of results to skip (between 0 and 900).")),
		mcp.WithNumber("number", mcp.Description("The maximum number of items to return (between 1 and 100). Defaults to 10.")),
		mcp.WithBoolean("limitLicense", mcp.Description("Whether the recipes should have an open license that allows display with proper attribution.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SearchrecipesHandler(cfg),
	}
}
