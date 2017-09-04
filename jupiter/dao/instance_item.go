package dao

import (
	"time"
	"weibo.com/opendcp/jupiter/models"
)

func InsertItem(item models.InstanceItem) error {
	o := GetOrmer()
	_, err := o.Insert(&item)
	return err
}

func InsertItems(items []models.InstanceItem) error {
	o := GetOrmer()
	/*err := o.Begin()
	if err != nil {
		return err
	}*/

	var err error
	for index, task := range items {
		_, err = o.Insert(&task)
		if err != nil {
			break
		}
		items[index].Id = task.Id
	}

	/*if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
	}*/
	return err
}

func UpdateItem(item models.InstanceItem) error {
	o := GetOrmer()
	/*err := o.Begin()
	if err != nil {
		return err
	}*/

	_, err := o.Update(&item)

	/*if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
	}*/
	return err

}

func DeleteItem(item models.InstanceItem) error {
	o := GetOrmer()
	_, err := o.Delete(&item)
	return err
}

func DeleteOldItems(before time.Time) error {
	o := GetOrmer()
	sql := "DELETE FROM task_item WHERE create_time < ?"
	_, err := o.Raw(sql, before).Exec()
	return err
}

func GetItemsByStatus(status int) ([]models.InstanceItem, error) {
	o := GetOrmer()
	var items []models.InstanceItem
	_, err := o.QueryTable(INSTANCE_ITEM_TABLE).RelatedSel().Filter("status", status).OrderBy("-id").All(&items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func GetItemById(id int) (models.InstanceItem, error) {
	o := GetOrmer()
	item := models.InstanceItem{Id: id}
	err := o.Read(&item)
	return item, err
}