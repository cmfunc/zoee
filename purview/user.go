package purview

// 用户权限

// User4PublicLesson 公开课的用户
type User4PublicLesson struct {
	*Parent4PublicLesson                      // 公开课家长
	Childs               []Child4PublicLesson // 家长的孩子
}

// Parent4PublicLesson 公开课家长
type Parent4PublicLesson struct {
	UserID uint32
}

// Child4PublicLesson 公开课家长的孩子
type Child4PublicLesson struct {
	RegularLesson struct {
		OrgType interface{} // 机构类型
		Company uint32      // 分公司
		Campus  uint32      // 校区
		Stage   uint8       // 学生所在阶段
	} // 孩子常规课信息
	QualityLesson struct {
		ApplyStage uint8 // 阶段
	} // 精品课
	IntalLearn struct{} // 国际游学
	Steam      struct{} // Steam
}
