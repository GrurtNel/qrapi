package category

func GetCategories() ([]*Category, error) {
	var categories []*Category
	err := categoryTable.FindAll(&categories)
	return categories, err
}
