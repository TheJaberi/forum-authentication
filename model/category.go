package forum

import (
	"context"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// TABLE: Category

func CreateCategory(name string) error {
	err := CreateCategoryDb(name)
	if err != nil {
		log.Println(NewCategoryError.Error() + err.Error())
		return err
	}
	newCategory, err := GetCategory(name)
	if err != nil {
		log.Println(NewCategoryError.Error() + err.Error())
		return err
	}
	AllCategories = append(AllCategories, newCategory)
	return nil
}

func CreateCategoryDb(name string) error {
	query := "INSERT INTO `Category` (`Name`) VALUES (?)"
	_, err := DB.ExecContext(context.Background(), query, name)
	if err != nil { // the category is added using the ExecContext
		return err
	}
	return nil
}

func GetCategories() error {
	AllCategories = nil
	var category Category
	categoryData, err := DB.Query("Select id, Name from Category")
	if err != nil {
		return err
	}
	defer categoryData.Close()
	for categoryData.Next() {
		err := categoryData.Scan(&category.CategoryID, &category.CategoryName)
		if err != nil {
			log.Println(ScanCategoryError.Error())
			return err
		}
		AllCategories = append(AllCategories, category)
	}
	return nil
}

func GetCategory(name string) (Category, error) {
	row := DB.QueryRow("SELECT id, Name from Category WHERE name=?", name)
	var category Category
	err := row.Scan(&category.CategoryID, &category.CategoryName)
	if err != nil {
		log.Println(ScanCategoryError.Error())
		return category, err
	}
	return category, nil
}

// TABLE: Post2Category

func AssignPostCategoryDb(postID int, postCategories []int) error {
	for _, category := range postCategories {
		queryCategory := "INSERT INTO `Post2Category` (`post_id`, `category_id`) VALUES (?, ?)"
		_, err := DB.ExecContext(context.Background(), queryCategory, postID, category)
		if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
			log.Println(err)
			return err
		}
	}
	err := GetCategories()
	if err != nil {
		return err
	}
	return nil
}

// gets the category id from the post id then the category name from categories
func GetPostCategories(p *Post) error {
	categoryData, err := DB.Query("Select category_id from Post2Category where post_id = ?", p.PostID)
	if err != nil {
		return errors.New("Category Query Error:" + err.Error())
	}
	defer categoryData.Close()
	for categoryData.Next() {
		var categoryID int
		err := categoryData.Scan(&categoryID)
		if err != nil {
			return errors.New("Category Scan Error:" + err.Error())
		}
		for i := range AllCategories {
			if categoryID == AllCategories[i].CategoryID {
				p.Category = append(p.Category, AllCategories[i])
				break
			}
		}
	}
	return nil
}
