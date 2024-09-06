package dao

import (
	"TTMS/kitex_gen/order"
	"context"
	"database/sql"
	"log"
	"strings"
)

func AddOrder(UserId, ScheduleId, SeatRow, SeatCol int, Date string, Price int) error {
	o := order.Order{UserId: int64(UserId), ScheduleId: int64(ScheduleId), SeatRow: int32(SeatRow), SeatCol: int32(SeatCol), Date: Date, Value: int32(Price), Type: 2}
	return DB.Create(&o).Error
}
func UpdateOrder(UserId, ScheduleId int64, SeatRow, SeatCol int32, Type int32, Date string) error {
	tx := DB.Begin()
	o := order.Order{}
	tx.Model(&o).Select("id").Where("user_id = ? and schedule_id = ? and seat_row = ? and seat_col = ?",
		UserId, ScheduleId, SeatRow, SeatCol).Order("id desc").Find(&o.Id)

	err := tx.Model(&order.Order{}).Where("id = ? ", o.Id).UpdateColumns(map[string]interface{}{
		"date": Date,
		"type": Type,
	}).Error
	tx.Commit()
	return err
}

func GetAllOrder(ctx context.Context, UserId, OrderType int) (orders []*order.Order, total int64, err error) {
	if OrderType == 2 {
		err = DB.WithContext(ctx).Where("user_id = ? and type = ?", UserId, OrderType).Order("date desc").Find(&orders).Error
		DB.WithContext(ctx).Where("user_id = ? and type = ?", UserId, OrderType).Order("date desc").Count(&total)
	} else {
		err = DB.WithContext(ctx).Where("user_id = ? and type != 2", UserId).Order("date desc").Find(&orders).Error
		DB.WithContext(ctx).Model(&order.Order{}).Where("user_id = ? and type != 2", UserId).Order("date desc").Count(&total)
	}
	return
}

func GetOrderAnalysis(ctx context.Context, ScheduleIdList []int64) (int64, int64, error) { //count,sum
	rows, err := DB.WithContext(ctx).Model(&order.Order{}).Select("count(user_id),coalesce(sum(value),0)").Where("schedule_id in ?", ScheduleIdList).Rows()
	if err != nil && strings.EqualFold(err.Error(), sql.ErrNoRows.Error()) {
		log.Println(err)
		log.Println(sql.ErrNoRows)
		return 0, 0, nil
	}
	rows.Next()
	var ans1, ans2 int64
	err = rows.Scan(&ans1, &ans2)
	if err != nil {
		log.Println(err)
	}
	return ans1, ans2, err
}
func DeleteOrder(ctx context.Context, userId, scheduleId, seatRow, seatCol int) error {
	o := order.Order{}
	err := DB.WithContext(ctx).Where("user_id = ? and schedule_id = ? and seat_row = ? and seat_col = ?", userId, scheduleId, seatRow, seatCol).Delete(&o).Error
	log.Println(&o)
	return err
}
func GetOrderCount(ctx context.Context, scheduleId, seatRow, seatCol int, types []int) int64 {
	var count int64 = 0
	err := DB.WithContext(ctx).Model(&order.Order{}).Where("schedule_id = ? and seat_row = ? and seat_col = ? and type in (?)", scheduleId, seatRow, seatCol, types).Count(&count).Error
	if err != nil {
		log.Println(ctx, "GetOrderCount err =", err)
	}
	return count
}
