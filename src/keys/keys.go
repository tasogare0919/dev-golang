package keys

import (
	"github.com/ChimeraCoder/anaconda"
)

func GetTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey("SrpLlbU8IgusU1w3Rx3WyOWKr")
	anaconda.SetConsumerSecret("En2TBTExQWLzpB6JJqhX5D63WzssfKLrhzsPttcXECJnkdKW9G")
	api := anaconda.NewTwitterApi("2901285342-ECiO57LwjlaQYqsJRDkvhidFLE63iBoJNrmjeKv", "En6tStauQFI5WbazlTVlnET3Vi5Au8Okc3jb693XGxasb")
	return api
}
