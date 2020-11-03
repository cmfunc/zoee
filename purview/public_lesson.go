package purview

//PublicLessonPurview 公开课分发范围
type PublicLessonPurview struct {
	*PublicLessonPurviewGoodsInfo
	*PublicLessonPurviewDistributeInfo
}

// PublicLessonPurviewGoodsInfo 商品信息
type PublicLessonPurviewGoodsInfo struct {
	ChargeType uint8
}

// PublicLessonPurviewDistributeInfo 公开课分发信息
type PublicLessonPurviewDistributeInfo struct {
	ApplayUsers       uint8 // all、riseReading、
	AppointUpperLimit uint64
	ViewRange         []PublicLessonPurviewDistributeInfoViewRangeItem // 可见范围：常规课、精品小班、国际游学、steam

}

// PublicLessonPurviewDistributeInfoViewRangeItem 可见范围类型
type PublicLessonPurviewDistributeInfoViewRangeItem interface {
	IsInRange() (bool, error)
}

// RegularLesson 常规课
type RegularLesson struct {
	OrgType    interface{} // 机构类型
	Company    uint32      // 分公司
	Campus     uint32      // 校区
	ApplyStage []uint8     // 适用阶段
}

// IsInRange 范围判断
func (r *RegularLesson) IsInRange() (bool, error) {
	return false, nil
}

// QualityLesson 精品课
type QualityLesson struct {
	ApplyStage []uint8 // 适用阶段
}

// IsInRange 范围判断
func (r *QualityLesson) IsInRange() (bool, error) {
	return false, nil
}

// IntalLearn 国际游学
type IntalLearn struct{}

// IsInRange 范围判断
func (r *IntalLearn) IsInRange() (bool, error) {
	return false, nil
}

// Steam Steam
type Steam struct{}

// IsInRange 范围判断
func (r *Steam) IsInRange() (bool, error) {
	return false, nil
}