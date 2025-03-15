package messages

// - /pantry - Manage your saved ingredient list

const WelcomeText = `* 🍳 Welcome to RecipePal! 🥗*

Hello %s! I'm RecipePal, your personal recipe assistant built to turn your available ingredients into delicious meals. No more staring at your fridge wondering what to cook!

* 🌟 What I Can Do For You *

- *Find recipes based on ingredients you have*
- *Suggest meals that match your dietary preferences*
- *Save your favorite recipes for later*
- *Help reduce food waste by using what you already have*
- *Provide step-by-step cooking instructions*

* 🚀 Getting Started *

Here are some commands to help you navigate:

- /start - Show this welcome message
- /findrecipe - Search recipes with your ingredients
- /diet - Set dietary preferences (vegetarian, gluten-free, etc.)
- /mealplan - Generate a weekly meal plan
- /help - Get detailed instructions

* 💡 Quick Tips *

- Separate multiple ingredients with commas
- Add cooking time limits with "quick" or "under 30 minutes"
- Use "breakfast," "lunch," "dinner," or "dessert" to specify meal type

* 📱 Try These Now *

👉 /findrecipe - Create your first recipe search
👉 /diet - Set your dietary preferences
👉 /help - Get more detailed instructions

I'm excited to help you discover new culinary creations! What would you like to cook today?`

const DietaryPreferenceButtonText = `*Dietary Preferences*`

const SetDietaryPreference = `*Dietary Preferences* 🥦

Tell me your dietary preferences by listing them separated by commas:

Example: vegetarian, gluten-free, dairy-free

Your preferences help me find recipes that match your needs!`

const SomethingWentWrong = `**Oops!** It looks like something went wrong on our end. We couldn't process your request at the moment. Please try again later. Thank you for your patience!`

const DietaryPreferencesSaved = `Great! We've saved your dietary preferences! 🍽️🥦 

Feel free to update them anytime! 😊`

const DietaryPreferenceSavedWithoutInvalid = `🎉 Great! We've saved your dietary preferences! 🍽️  

However, we noticed a few options that were invalid (%s) and couldn't be saved. 🚫  

Feel free to update them anytime! 😊  `

const InvalidDietaryPreference = `We noticed a few options that were invalid (%s) and couldn't be saved. 🚫  `

const InvalidCallback = `Invalid callback data`

const FailedToGetResponse = `Failed to get response`

const DietaryPreferenceDeleted = `Dietary preference deleted successfully!`

const FindRecipes = `🎉 **Let's Cook Something Delicious!** 🥢

We're excited 🌟 to help you whip up something amazing in the kitchen. 🍳

Please list the ingredients you have right now, and we'll generate a fantastic recipe just for you! 🥩

📝 *How to list your ingredients:*
- Separate each ingredient with a comma.
- Include quantities if you can (e.g., 2 eggs, 1 cup flour).
- Don't worry about being too precise; we'll do our best to match your ingredients to a great recipe! 🥩

Example: 2 eggs, 1 cup flour, 1/2 cup milk, 1 tsp sugar

Ready? Send us your ingredients, and let's get cooking! 🥩

`
