package theme

type ThemeStatus byte

const (
	ThemeStatusPending    ThemeStatus = iota //正在准备，仅供预览
	ThemeStatusOpening                       //开放中
	ThemeStatusDeprecated                    //曾经存在过，如今弃用
	ThemeStatusDeleted                       //手动删除
)

type Theme struct {
	ID           int64
	Title        string //主题名称
	Desc         string //主题说明
	MinMember    int    //最小参与人数
	MaxMember    int    //最大参与人数
	PlayDuration int    //规定游玩时长 seconds
	Status       ThemeStatus
	Tips         []*Tip       `orm:"reverse(many)"`
	TimeRange    []*TimeRange `orm:"reverse(many)"`
}

func (t *Theme) TableName() string {
	return "theme"
}
