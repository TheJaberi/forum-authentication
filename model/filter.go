package forum

import (
	"errors"
	"strconv"
)

// filters depending on category
func FilterByCategory(categoryID string) error {
	var filteredPosts []Post
	category, err := strconv.Atoi(categoryID)
	if err != nil {
		return err
	}
	if category > len(AllData.AllCategories) {
		return errors.New("Category not found")
	}
	for i := 0; i < len(AllPosts); i++ {
		for j := 0; j < len(AllPosts[i].Category); j++ {
			if category == AllPosts[i].Category[j].CategoryID {
				filteredPosts = append(filteredPosts, AllPosts[i])
				break
			}
		}
	}
	AllData.AllPosts = RSort(filteredPosts)
	AllData.CategoryCheck = false
	AllData.LoggedUser = LoggedUser
	return nil
}

// filters depending on the user data
func FilterUserData(path string) error {
	var filteredPosts []Post
	userID := AllData.LoggedUserID
	for i := 0; i < len(AllPosts); i++ {
		if path == "myposts" && AllPosts[i].UserID == userID {
			filteredPosts = append(filteredPosts, AllPosts[i])
		}
		if path == "mylikes" && AllPosts[i].Userlike {
			filteredPosts = append(filteredPosts, AllPosts[i])
		}
		if path == "mydislikes" && AllPosts[i].UserDislike {
			filteredPosts = append(filteredPosts, AllPosts[i])
		}
	}
	AllData.AllPosts = RSort(filteredPosts)
	// AllData.CategoryCheck = false
	return nil
}
