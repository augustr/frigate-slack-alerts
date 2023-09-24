package sharedTypes

type EventBody struct {
	ID              string   `json:"id"`
	Camera          string   `json:"camera"`
	FrameTime       float64  `json:"frame_time"`
	SnapshotTime    float64  `json:"snapshot_time"`
	Label           string   `json:"label"`
	SubLabel        *string  `json:"sub_label"`
	TopScore        float64  `json:"top_score"`
	FalsePositive   bool     `json:"false_positive"`
	StartTime       float64  `json:"start_time"`
	EndTime         *float64 `json:"end_time"`
	Score           float64  `json:"score"`
	Box             []int    `json:"box"`
	Area            int      `json:"area"`
	Ratio           float64  `json:"ratio"`
	Region          []int    `json:"region"`
	Stationary      bool     `json:"stationary"`
	MotionlessCount float64  `json:"motionless_count"`
	PositionChanges int      `json:"position_changes"`
	CurrentZones    []string `json:"current_zones"`
	HasClip         bool     `json:"has_clip"`
	HasSnapshot     bool     `json:"has_snapshot"`
}

type EventMessage struct {
	Type   string    `json:"type"`
	Before EventBody `json:"before"`
	After  EventBody `json:"after"`
}
