package fitbit

type FoodService service

type FoodLog struct {
	Foods   []FoodLogEntry    `json:"food"`
	Summary NutritionalValues `json:"summary"`
	Goals   FoodGoals         `json:"goals"`
}

type FoodLogEntry struct {
	IsFavourite       bool              `json:"isFavourite"`
	LogDate           string            `json:"logDate"`
	LogId             int               `json:"logId"`
	LoggedFood        Food              `json:"loggedFood"`
	NutritionalValues NutritionalValues `json:"nutritionalValues"`
}

type Food struct {
	AccessLevel string  `json:"accessLevel"`
	Amount      float64 `json:"amount"`
	Brand       string  `json:"brand"`
	Calories    Calorie `json:"calories"`
	FoodId      int     `json:"foodId"`
	MealTypeId  int     `json:"mealTypeId"`
	Locale      string  `json:"locale"`
	Name        string  `json:"name"`
	Unit        Unit    `json:"unit"`
	Units       []int   `json:"units"`
}

type Unit struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Plural string `json:"plural"`
}

type NutritionalValues struct {
	Calories Calorie `json:"calories"`
	Carbs    float64 `json:"carbs"`
	Fat      float64 `json:"fat"`
	Fiber    float64 `json:"fiber"`
	Protein  float64 `json:"protein"`
	Sodium   float64 `json:"sodium"`
	Water    float64 `json:"water"`
}

type FoodGoals struct {
	Calories Calorie `json:"calories"`
}

type Meals struct {
	Meals []Meal `json:"meals"`
}

type Meal struct {
	MealFoods   []Food `json:"mealFoods"`
	Id          int    `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}
