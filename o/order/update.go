package order

import (
	"qrapi/g/x/web"
)

func DeleteByID(id string) error {
	var order, err = GetOrderByID(id)
	if err != nil || order != nil {
		return web.BadRequest("Không tồn tại sản phẩm")
	}
	if order.Activated {
		return web.BadRequest("Sản phẩm đang phân phối nên không thể xóa")
	}
	return orderTable.DeleteByID(id)
}
