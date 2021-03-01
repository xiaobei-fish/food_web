package models

import (
	"fmt"
	"food_web/utils"
)

//删除食品
func DeleteFoodWithId(foodID string) (int64, error) {
	sql:= "delete from food_info where id=" + foodID
	fmt.Println(sql)
	return utils.ModifyDB(sql)
}

//删除用户
func DeleteUserWithId(userID string) (int64, error) {
	sql:= "delete from users where id=" + userID
	fmt.Println(sql)
	return utils.ModifyDB(sql)
}
