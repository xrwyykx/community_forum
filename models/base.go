package models

type Page struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

//type Time time.Time //定义了time.Time的别名，也是一个新的类型，不会影响到其它使用time.Time的代码
//
//// 实现时间序列化
//func (t Time) MarshalJSON() ([]byte, error) {
//	return json.Marshal(time.Time(t).Format("2006-01-02 15:04:05")) //将t转化为time.Time类型，然后将其序列化为后面的格式
//}
//func (t *Time) UnmarshalJSON(data []byte) error { //调用函数的对象是Time类型的指针，data是函数调用时提供的参数，error是返回值的类型
//	var str string
//	if err := json.Unmarshal(data, &str); err != nil { //将该序列化对象反序列化成字符串保存在str中
//		return err
//	}
//	tmp, err := time.Parse("2006-01-02 15:04:05", str) //将字符串转化为对应的格式，然后保存在tmp中
//	if err != nil {
//		return err
//	}
//	*t = Time(tmp) //将t的值赋值为转化为Time类型的tmp
//	return nil
//}
