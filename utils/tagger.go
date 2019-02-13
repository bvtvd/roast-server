package utils

import (
	. "roast-server/models"
	"strings"
	. "roast-server/database"
)

func TagCafe(cafeId int, tags []string, userId int) {
	var tagTemp Tag

	for _, tag := range tags {
		tagTemp = make(Tag{})

		name := strings.TrimSpace(tag);

		// 获取标签实例
		Db.Where(Tag{Name: name}).FirstOrCreate(&tagTemp)

		var cafesUsersTag CafesUsersTag = CafesUsersTag{CafeId: cafeId, UserId: userId, TagId: tagTemp.id}

		// 创建关联关系
		Db.FirstOrCreate(&cafesUsersTag)
	}
}