package v1

//func GetConnectorByID(c *gin.Context){
//	id,_ := strconv.Atoi(c.Param("id"))
//	conn := m.OraConnectors{Id:id}
//	conn.GetConnectorByID()
//
//	c.JSON(http.StatusOK, gin.H{
//		"code":   0,
//		"status": "ok",
//		"data":  conn,
//	})
//
//}
//
//func GetConnectorByIP(c *gin.Context){
//	ipaddr:="10"
//	rt := m.GetConnectorByIP(ipaddr)
//
//	c.JSON(http.StatusOK, gin.H{
//		"code":   0,
//		"status": "ok",
//		"data":  rt,
//	})
//}
//
//type OraHang struct{
//	Name string
//	Waiters int
//	Holders int
//}
//func GetHangs(c *gin.Context){
//	fmt.Println("DBH=", DBH)
//	var oh []OraHang
//
//	for _,v:=range DBH {
//		var o OraHang
//		o.Name = v.DBS.Name
//		o.Waiters = v.Waiters
//		o.Holders = v.Holders
//		oh = append(oh,o)
//	}
//
//	c.JSON(http.StatusOK,gin.H{
//		"result":oh,
//	})
//}