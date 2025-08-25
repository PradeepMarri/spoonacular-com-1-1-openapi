package main

import (
	"github.com/spoonacular-api/mcp-server/config"
	"github.com/spoonacular-api/mcp-server/models"
	tools_meal_planning "github.com/spoonacular-api/mcp-server/tools/meal_planning"
	tools_wine "github.com/spoonacular-api/mcp-server/tools/wine"
	tools_ingredients "github.com/spoonacular-api/mcp-server/tools/ingredients"
	tools_products "github.com/spoonacular-api/mcp-server/tools/products"
	tools_misc "github.com/spoonacular-api/mcp-server/tools/misc"
	tools_recipes "github.com/spoonacular-api/mcp-server/tools/recipes"
	tools_menu_items "github.com/spoonacular-api/mcp-server/tools/menu_items"
	tools_food "github.com/spoonacular-api/mcp-server/tools/food"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_meal_planning.CreateGetshoppinglistTool(cfg),
		tools_wine.CreateGetdishpairingforwineTool(cfg),
		tools_wine.CreateGetwinerecommendationTool(cfg),
		tools_wine.CreateGetwinepairingTool(cfg),
		tools_ingredients.CreateGetingredientsubstitutesTool(cfg),
		tools_products.CreateSearchgroceryproductsbyupcTool(cfg),
		tools_products.CreateAutocompleteproductsearchTool(cfg),
		tools_meal_planning.CreateAddtomealplanTool(cfg),
		tools_ingredients.CreateMapingredientstogroceryproductsTool(cfg),
		tools_misc.CreateGetrandomfoodtriviaTool(cfg),
		tools_misc.CreateSearchallfoodTool(cfg),
		tools_products.CreateSearchgroceryproductsTool(cfg),
		tools_recipes.CreateGetrecipeequipmentbyidTool(cfg),
		tools_recipes.CreateGetrecipeinformationbulkTool(cfg),
		tools_meal_planning.CreateConnectuserTool(cfg),
		tools_products.CreateClassifygroceryproductTool(cfg),
		tools_misc.CreateSearchcustomfoodsTool(cfg),
		tools_ingredients.CreateIngredientsearchTool(cfg),
		tools_products.CreateGetcomparableproductsTool(cfg),
		tools_menu_items.CreateSearchmenuitemsTool(cfg),
		tools_meal_planning.CreateAddtoshoppinglistTool(cfg),
		tools_meal_planning.CreateGetmealplantemplateTool(cfg),
		tools_recipes.CreateGetrecipeingredientsbyidTool(cfg),
		tools_recipes.CreateGetrecipetastebyidTool(cfg),
		tools_recipes.CreateSearchrecipesTool(cfg),
		tools_recipes.CreateGuessnutritionbydishnameTool(cfg),
		tools_misc.CreateImageclassificationbyurlTool(cfg),
		tools_recipes.CreateComputeglycemicloadTool(cfg),
		tools_recipes.CreateGetrecipenutritionwidgetbyidTool(cfg),
		tools_recipes.CreateQuickanswerTool(cfg),
		tools_meal_planning.CreateGetmealplanweekTool(cfg),
		tools_recipes.CreateGetanalyzedrecipeinstructionsTool(cfg),
		tools_meal_planning.CreateGetmealplantemplatesTool(cfg),
		tools_products.CreateClassifygroceryproductbulkTool(cfg),
		tools_meal_planning.CreateGeneratemealplanTool(cfg),
		tools_ingredients.CreateGetingredientinformationTool(cfg),
		tools_wine.CreateGetwinedescriptionTool(cfg),
		tools_recipes.CreateConvertamountsTool(cfg),
		tools_misc.CreateGetarandomfoodjokeTool(cfg),
		tools_recipes.CreateGetsimilarrecipesTool(cfg),
		tools_products.CreateGetproductinformationTool(cfg),
		tools_recipes.CreateCreaterecipecardgetTool(cfg),
		tools_misc.CreateTalktochatbotTool(cfg),
		tools_recipes.CreateGetrecipepricebreakdownbyidTool(cfg),
		tools_recipes.CreateGetrandomrecipesTool(cfg),
		tools_misc.CreateSearchsitecontentTool(cfg),
		tools_recipes.CreateSearchrecipesbynutrientsTool(cfg),
		tools_misc.CreateImageanalysisbyurlTool(cfg),
		tools_ingredients.CreateComputeingredientamountTool(cfg),
		tools_misc.CreateGetconversationsuggestsTool(cfg),
		tools_menu_items.CreateAutocompletemenuitemsearchTool(cfg),
		tools_food.CreateSearchrestaurantsTool(cfg),
		tools_recipes.CreateAnalyzerecipeTool(cfg),
		tools_menu_items.CreateGetmenuiteminformationTool(cfg),
		tools_recipes.CreateSummarizerecipeTool(cfg),
		tools_ingredients.CreateGetingredientsubstitutesbyidTool(cfg),
		tools_recipes.CreateGetrecipeinformationTool(cfg),
		tools_misc.CreateSearchfoodvideosTool(cfg),
		tools_recipes.CreateExtractrecipefromwebsiteTool(cfg),
		tools_recipes.CreateAnalyzearecipesearchqueryTool(cfg),
		tools_recipes.CreateSearchrecipesbyingredientsTool(cfg),
		tools_ingredients.CreateAutocompleteingredientsearchTool(cfg),
		tools_recipes.CreateAutocompleterecipesearchTool(cfg),
	}
}
