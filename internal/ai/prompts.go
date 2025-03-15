package ai

const FIND_RECIPE_PROMPT = `Generate 5 personalized recipe suggestions based on user-specified available ingredients and dietary preferences. 

The user will provide:
- A list of available ingredients: {userIngredients}
- Their dietary preference: {userDietaryPreference} (e.g., vegan, gluten-free, low-carb, etc.)

Your task is to:
1. Use as many of the provided ingredients as possible.
2. Ensure the recipe adheres to the user's dietary preferences.
3. Provide the output as an array of objects in the following JSON format:

[
    {
        "title": "🍴 [Recipe Name]",
        "ingredients": "🥕 Ingredients:\n- [Ingredient 1]\n- [Ingredient 2]\n- [Ingredient 3]",
        "instructions": "📝 Instructions:\n1. [Step 1]\n2. [Step 2]\n3. [Step 3]",
        "dietaryCompliance": "🌱 Dietary Compliance:\n- [How the recipe meets the dietary preferences]"
    }
]
Rules:
Always prioritize the user's provided ingredients and dietary preferences.

Use emojis to make the output more engaging and user-friendly. 🎉

Keep the instructions simple and easy to follow.

Ensure the recipes are creative and practical.

Example Input:
ingredients:{userIngredients}
dietaryPreference:{userDietaryPreference}

Example Output:
[
    {
        "title": "🍴 Garlic Spinach Stuffed Chicken Breast",
        "ingredients": "🥕 Ingredients:\n- Chicken breast\n- Spinach\n- Garlic\n- Olive oil\n- Salt\n- Pepper",
        "instructions": "📝 Instructions:\n1. Preheat the oven to 375°F (190°C).\n2. Sauté spinach and garlic in olive oil until wilted.\n3. Stuff the chicken breast with the spinach mixture.\n4. Bake for 25 minutes or until the chicken is fully cooked.",
        "dietaryCompliance": "🌱 Dietary Compliance:\n- This recipe is low-carb and high in protein. Perfect for a low-carb diet! 🥗"
    }
]

Notes:
If an ingredient cannot be used due to dietary restrictions, exclude it and suggest alternatives if possible.

Be creative with the recipes while keeping them practical and easy to prepare.
`
