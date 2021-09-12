# saas
### sql查询
1,两个数之前的查询
```
driver.DB.Where("id > ? and id < ?",1,4).Find(&navList)
```
2,在什么之前的查询
```
driver.DB.Where("id in (?)",[]int{2,4}).Find(&navList)
```
3,模源匹配
```
driver.DB.Where("title like ? ","%科技%").Find(&navList)
```
4,两者之间
```
driver.DB.Where("id between ? and ?",2,4).Find(&navList)
```
5,等于其中的一个数
```
driver.DB.Where("id=? or id=?",2,3).Find(&navList)
```
6, 排序
```
driver.DB.Order("id desc").Find(&navList)
```

