package dao

import (
	"TTMS/kitex_gen/studio"
	"context"
	"errors"
	"gorm.io/gorm"
)

func AddBatchSeat(ctx context.Context, studioInfo *Studio) error {
	id, row, col := int(studioInfo.Id), int(studioInfo.RowsCount), int(studioInfo.ColsCount)
	seats := make([]*studio.Seat, 0, row*col)
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			seats = append(seats, &studio.Seat{StudioId: int64(id), Row: int64(i), Col: int64(j), Status: 1})
		}
	}
	return DB.WithContext(ctx).Create(seats).Error
}
func AddSeat(ctx context.Context, seatInfo *studio.Seat) error {
	s1 := studio.Studio{Id: seatInfo.StudioId}
	DB.WithContext(ctx).Find(&s1)
	if s1.Id == 0 {
		return errors.New("演出厅不存在")
	}
	if seatInfo.Row > s1.RowsCount || seatInfo.Col > s1.ColsCount {
		return errors.New("无法添加超出演出厅规模的座位")
	}
	s := studio.Seat{}
	if DB.WithContext(ctx).Where("studio_id = ? and row = ? and col = ? ",
		seatInfo.StudioId, seatInfo.Row, seatInfo.Col).
		Limit(1).Find(&s); s.Id > 0 && s.Status > 0 {
		return errors.New("不允许的行为：同一位置重复加入座位")
	} else if s.Id == 0 {
		//没有该座位的数据存在
		return DB.WithContext(ctx).Create(seatInfo).Error
	} else {
		//该座位的数据存在，只是status==0
		return UpdateSeat(ctx, seatInfo)
	}

}
func GetAllSeat(ctx context.Context, studioId, Current, PageSize int) ([]*studio.Seat, int64, error) {
	var seats []*studio.Seat
	tx := DB.WithContext(ctx).Where("studio_id = ?", studioId).Order("row").
		Order("col").Offset((Current - 1) * PageSize).Limit(PageSize).Find(&seats)
	var total int64
	tx = DB.WithContext(ctx).Model(&studio.Seat{}).Where("studio_id = ?", studioId).Order("row").
		Order("col").Count(&total)
	return seats, total, tx.Error
}
func UpdateSeat(ctx context.Context, seatInfo *studio.Seat) error {
	s := studio.Seat{}
	if DB.WithContext(ctx).Where("studio_id = ? and row = ? and col = ? ", seatInfo.StudioId,
		seatInfo.Row, seatInfo.Col).Limit(1).Find(&s); s.Id > 0 {
		return DB.WithContext(ctx).Model(&s).Where("studio_id = ? and row = ? and col = ? ",
			seatInfo.StudioId, seatInfo.Row, seatInfo.Col).Update("status", seatInfo.Status).Error
	}
	if seatInfo.Status == 0 { //删除了一个座位
		DB.WithContext(ctx).Model(&studio.Studio{}).Where("id = ?", seatInfo.StudioId).
			UpdateColumn("seats_count", gorm.Expr("seats_count - 1"))
	}
	return errors.New("该位置上无座位，修改失败")
}
func DeleteSeat(ctx context.Context, seatInfo *studio.Seat) error {
	seatInfo.Status = 0
	return UpdateSeat(ctx, seatInfo)
}
func RealDeleteSeat(studioId, row, col int) error {
	return DB.Where("studio_id=? and row=? and col=?", studioId, row, col).Delete(&studio.Seat{}).Error
}
