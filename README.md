# saas
### sql查询
1,两个数之前的查询
```
driver.DB.Where("id > ? and id < ?",1,4).Find(&navList)
```
