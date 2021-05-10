package utils

/*
	Struct for the weartherAPI response
*/
type WeatherStruct struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type Fridge struct {
	Ingredients []string
}

//Recipe struct holds all data from the spoonacular API used to find recipes
type Recipe []struct {
	Name                 string           `json:"title"`
	Image                string           `json:"image"`
	MissedIngredients    []IngredientList `json:"missedIngredients"`   //Missing ingredients from
	UsedIngredientsCount int              `json:"usedIngredientCount"` //Amounts of Ingredients used
	UsedIngredients      []IngredientList `json:"usedIngredients"`
}

type IngredientList struct {
	IngredientName string `json:"name"`
}

// NewsAnswer This struct formats the answer provided by the Bing News Search API.
type NewsAnswer struct {
	ReadLink       string `json:"readLink"`
	QueryContext   struct {
		OriginalQuery   string   `json:"originalQuery"`
		AdultIntent     bool        `json:"adultIntent"`
	} `json:"queryContext"`
	TotalEstimatedMatches   int  `json:"totalEstimatedMatches"`
	Sort  []struct {
		Name    string  `json:"name"`
		ID       string    `json:"id"`
		IsSelected       bool  `json:"isSelected"`
		URL      string   `json:"url"`
	}  `json:"sort"`
	Value   []struct   {
		Name     string   `json:"name"`
		URL   string    `json:"url"`
		Image   struct   {
			Thumbnail   struct  {
				ContentUrl  string  `json:"thumbnail"`
				Width   int  `json:"width"`
				Height  int   `json:"height"`
			} `json:"thumbnail"`
		} `json:"image"`
		Description  string  `json:"description"`
		Provider  []struct   {
			Type   string    `json:"_type"`
			Name  string     `json:"name"`
		} `json:"provider"`
		DatePublished   string   `json:"datePublished"`
	} `json:"value"`
}
