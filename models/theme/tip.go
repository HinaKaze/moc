package theme

type Tip struct {
	ID    int64
	Theme *Theme `orm:"rel(fk)"`
	Stage string //阶段信息
	Desc  string //提示说明
}

func (t *Tip) TableName() string {
	return "theme_tip"
}
