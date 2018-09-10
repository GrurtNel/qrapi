package product

func DeleteByID(id string) error {
	return productTable.DeleteByID(id)
}
