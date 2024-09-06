package dao

import (
	"TTMS/kitex_gen/studio"
	"context"
	"errors"
)

func AddStudio(ctx context.Context, Name string, row, col int64) error {
	s := Studio{Name: Name, RowsCount: row, ColsCount: col}
	tx := DB.WithContext(ctx).Create(&s)
	if tx.Error != nil {
		return tx.Error
	}
	return AddBatchSeat(ctx, &s)
}
func GetAllStudio(ctx context.Context, Current, PageSize int) ([]*studio.Studio, int64, error) {
	studios := make([]*studio.Studio, PageSize)
	tx := DB.WithContext(ctx).Select("studios.*,count(s.studio_id) as seats_count").
		Joins("join seats as s on studios.id=s.studio_id").Where("s.status=1").
		Group("studios.id").Offset((Current - 1) * PageSize).Limit(PageSize).Find(&studios)
	var total int64
	tx = DB.WithContext(ctx).Model(&Studio{}).Count(&total)
	return studios, total, tx.Error
}

func GetStudio(ctx context.Context, id int) (*studio.Studio, error) {
	var s studio.Studio
	tx := DB.WithContext(ctx).Table("studios as s1").Select("s1.*,count(s2.studio_id) as seats_count").
		Joins("join seats as s2 on s1.id=s2.studio_id").Where("s2.status=1 and s1.id = ?", id).
		Group("s2.studio_id").Limit(1).Find(&s)
	return &s, tx.Error
}

func UpdateStudio(ctx context.Context, StudioInfo *studio.Studio) error {
	s := studio.Studio{}
	DB.WithContext(ctx).Where("id = ?", StudioInfo.Id).Limit(1).Find(&s)
	if s.Id > 0 {
		if StudioInfo.RowsCount > s.RowsCount || StudioInfo.ColsCount > s.ColsCount {
			return errors.New("座位规模不能比原来更大")
		}
		if StudioInfo.RowsCount == 0 {
			StudioInfo.RowsCount = s.RowsCount
		}
		if StudioInfo.ColsCount == 0 {
			StudioInfo.ColsCount = s.ColsCount
		}
		tx := DB.Begin()
		if StudioInfo.RowsCount > 0 || StudioInfo.ColsCount > 0 { //演出厅规模变小之后删除座位
			var err error
			for i := 1; i <= int(s.RowsCount); i++ {
				for j := 1; j <= int(s.ColsCount); j++ {
					if i > int(StudioInfo.RowsCount) || j > int(StudioInfo.ColsCount) {
						err = RealDeleteSeat(int(StudioInfo.Id), i, j)
					}
				}
			}
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		err := DB.WithContext(ctx).Model(&studio.Studio{}).Where("id = ?", StudioInfo.Id).Updates(&StudioInfo).Error
		if err != nil { //修改成功
			tx.Rollback()
			return err
		}
		tx.Commit()
	} else {
		return errors.New("该演出厅不存在")
	}
	return nil
}
func DeleteStudio(ctx context.Context, id int64) error {
	s := studio.Studio{}
	DB.WithContext(ctx).Where("id  = ?", id).Limit(1).Find(&s)
	if s.Id > 0 {
		return DB.WithContext(ctx).Delete(&s).Error
	} else {
		return errors.New("该演出厅不存在")
	}
}
