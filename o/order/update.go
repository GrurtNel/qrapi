package order

import (
	"qrapi/g/x/web"
)

func DeleteByID(id string) error {
	var order, err = GetOrderByID(id)
	if err != nil || order == nil {
		return web.BadRequest("Không tồn tại đơn hàng")
	}
	if order.Activated {
		return web.BadRequest("Đơn hàng đang phân phối nên không thể xóa")
	}
	return orderTable.DeleteByID(id)
}
