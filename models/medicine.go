package models

type UsageRule string

const (
	SebelumMakan UsageRule = "sebelum_makan"
	SesudahMakan UsageRule = "sesudah_makan"
)

type Medicine struct {
	Model
	Name           string    `bson:"name"`
	TabletCount    int       `bson:"tablet_count"`
	UsageRule      UsageRule `bson:"usage_rule"`
	UsageFrequency int       `bson:"usage_frequency"`
	UsageTimes     []string  `bson:"usage_times"`
	IsActivated    bool      `bson:"is_activated"`
}
